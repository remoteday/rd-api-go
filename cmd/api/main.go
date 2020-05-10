package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/remoteday/rd-api-go/src/config"
	web "github.com/remoteday/rd-api-go/src/platform/adapters/web"
	_ "github.com/remoteday/rd-api-go/src/platform/docs"
	log "github.com/sirupsen/logrus"
)

var build = "develop"

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file --", err)
	}

	flag.Parse()

	var cfg struct {
		Web  config.Web
		DB   config.Database
		Auth config.Auth
	}

	if err := envconfig.Process("rd-api", &cfg); err != nil {
		log.Fatalf("main : Parsing Config : %v", err)
	}

	log.Infof("main : Started : Application Initializing version %q", build)
	defer log.Infoln("main : Completed")

	// Configure DB
	connection := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", cfg.DB.DBHost, cfg.DB.DBPort, cfg.DB.DBUser, cfg.DB.DBPass, cfg.DB.DBName)

	dbConn, err := sql.Open("postgres", connection)

	if err != nil {
		fmt.Println(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Start Debug Service

	// /debug/vars - Added to the default mux by the expvars package.
	// /debug/pprof - Added to the default mux by the net/http/pprof package.

	debug := http.Server{
		Addr:           cfg.Web.DebugHost + ":" + cfg.Web.DebugPort,
		Handler:        http.DefaultServeMux,
		ReadTimeout:    cfg.Web.ReadTimeout,
		WriteTimeout:   cfg.Web.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// Not concerned with shutting this down when the
	// application is being shutdown.
	go func() {
		log.Infof("main : Debug Listening %s", cfg.Web.DebugHost+":"+cfg.Web.DebugPort)
		log.Infof("main : Debug Listener closed : %v", debug.ListenAndServe())
	}()

	// Start API Service
	r := web.NewWebAdapter(dbConn, cfg.DB, cfg.Auth)

	api := http.Server{
		Addr:           cfg.Web.APIHost + ":" + cfg.Web.APIPort,
		Handler:        r,
		ReadTimeout:    cfg.Web.ReadTimeout,
		WriteTimeout:   cfg.Web.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		log.Infof("main : API Listening %s", cfg.Web.APIHost+":"+cfg.Web.APIPort)
		serverErrors <- api.ListenAndServe()
	}()

	// Shutdown

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	// =========================================================================
	// Stop API Service

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		log.Fatalf("main : Error starting server: %v", err)

	case <-osSignals:
		log.Infoln("main : Start shutdown...")

		// Create context for Shutdown call.
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		// Asking listener to shutdown and load shed.
		if err := api.Shutdown(ctx); err != nil {
			log.Infof("main : Graceful shutdown did not complete in %v : %v", cfg.Web.ShutdownTimeout, err)
			if err := api.Close(); err != nil {
				log.Fatalf("main : Could not stop http server: %v", err)
			}
		}
	}

}
