package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/remoteday/rd-api-go/src/middlewares/negronilogrus"
	"github.com/remoteday/rd-api-go/src/platform"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/urfave/negroni"
)

// Handler represent the httphandler for healthcheck
type Handler struct {
	App platform.App
}

// @title API
// @version 1.0
// @host 0.0.0.0:3000
// @BasePath /

// NewWebAdapter -
func NewWebAdapter(app platform.App) http.Handler {
	r := mux.NewRouter()

	// Redirects from `/docs` to `/docs/index.html`
	r.HandleFunc("/docs", func(res http.ResponseWriter, req *http.Request) {
		http.Redirect(res, req, fmt.Sprintf("%s/index.html", req.URL), http.StatusSeeOther)
	}).Methods("GET")
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	// App Routes
	NewHealthcheckHTTPHandler(r, app)

	loglevel := logrus.InfoLevel
	n := negroni.New()
	n.Use(negronilogrus.NewCustomMiddleware(loglevel, &logrus.JSONFormatter{}, "web"))
	n.UseHandler(r)
	return n
}
