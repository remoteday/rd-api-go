package platform

import (
	"database/sql"

	"go.uber.org/zap"
)

// App -
type App struct {
	DbConn *sql.DB
	Logger *zap.Logger
}

// NewApp -
func NewApp(dbConn *sql.DB, logger *zap.Logger) App {
	return App{
		DbConn: dbConn,
	}
}
