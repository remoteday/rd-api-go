package web

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/remoteday/rd-api-go/cmd/api/routes"
	"github.com/remoteday/rd-api-go/src/config"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/urfave/negroni"
)

// @title API
// @version 1.0
// @host 0.0.0.0:3000
// @BasePath /

// NewWebAdapter -
func NewWebAdapter(dbConn *sql.DB, dbConfig config.Database, authConfig config.Auth) http.Handler {
	r := mux.NewRouter()

	// Redirects from `/docs` to `/docs/index.html`
	r.HandleFunc("/docs", func(res http.ResponseWriter, req *http.Request) {
		http.Redirect(res, req, fmt.Sprintf("%s/index.html", req.URL), http.StatusSeeOther)
	}).Methods("GET")
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	// App Routes
	routes.NewHealthcheckHTTPHandler(r, dbConn)

	n := negroni.Classic()
	n.UseHandler(r)
	return n
}
