package routes

import (
	"database/sql"

	"github.com/remoteday/rd-api-go/src/config"
)

// Handler represent the httphandler for healthcheck
type Handler struct {
	dbConn     *sql.DB
	dbConfig   config.Database
	authConfig config.Auth
}
