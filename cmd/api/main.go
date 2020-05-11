package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/remoteday/rd-api-go/cmd/api/routes"
	"github.com/remoteday/rd-api-go/src/config"
	"github.com/remoteday/rd-api-go/src/db"
	_ "github.com/remoteday/rd-api-go/src/docs"
	log "github.com/sirupsen/logrus"
)

var build = "develop"

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file --", err)
	}

	flag.Parse()

	var cfg config.AppConfig

	if err := envconfig.Process("rd-api", &cfg); err != nil {
		log.Fatalf("main : Parsing Config : %v", err)
	}

	log.Infof("main : Started : Application Initializing version %q", build)
	defer log.Infoln("main : Completed")

	// Configure DB
	app, err := InitializeApp(cfg)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = app.DbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer func() {
		err := app.DbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Start seed
	switch flag.Arg(0) {
	case "with-migrate":
		if err := db.Migrate(app.DbConn); err != nil {
			log.Println("error applying migrations", err)
			os.Exit(1)
		}
		log.Println("Migrations complete")
	case "migrate":
		if err := db.Migrate(app.DbConn); err != nil {
			log.Println("error applying migrations", err)
			os.Exit(1)
		}
		log.Println("Migrations complete")
		return

	case "seed":
		if err := db.Seed(app.DbConn); err != nil {
			log.Println("error seeding database", err)
			os.Exit(1)
		}
		log.Println("Seed data complete")
		return

	}

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
	r := routes.NewWebAdapter(app)

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
		log.Infof("main : API Listening http://%s", cfg.Web.APIHost+":"+cfg.Web.APIPort)
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
