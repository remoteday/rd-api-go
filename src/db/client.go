package db

import (
	"database/sql"
	"fmt"

	"github.com/remoteday/rd-api-go/src/config"
)

// NewDatabaseConnection -
func NewDatabaseConnection(config config.AppConfig) (*sql.DB, error) {
	// Configure DB
	connection := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", config.DB.DBHost, config.DB.DBPort, config.DB.DBUser, config.DB.DBPass, config.DB.DBName)

	dbConn, err := sql.Open("postgres", connection)

	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
