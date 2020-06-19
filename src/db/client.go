package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/remoteday/rd-api-go/src/config"
)

// NewDatabaseConnection -
func NewDatabaseConnection(config config.AppConfig) (*sqlx.DB, error) {
	// Configure DB
	connection := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", config.DB.DBHost, config.DB.DBPort, config.DB.DBUser, config.DB.DBPass, config.DB.DBName)

	dbConn, err := sqlx.Connect("postgres", connection)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	return dbConn, nil
}
