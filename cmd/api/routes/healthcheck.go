package routes

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/remoteday/rd-api-go/src/healthcheck"
)

// NewHealthcheckHTTPHandler -
// @Summary Healthcheck
// @Description application health check
// @Accept  json
// @Produce  json
// @Success 200 {object} healthcheck.HealthCheck
// @Router /_health [get]
func NewHealthcheckHTTPHandler(r *mux.Router, dbConn *sql.DB) {
	handler := &Handler{}
	r.HandleFunc("/_health", handler.Healthcheck).Methods("GET")
	r.HandleFunc("/", handler.Healthcheck).Methods("GET")
}

// Healthcheck -
func (a *Handler) Healthcheck(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(healthcheck.HealthCheck{Status: "Ok"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, err := w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}
