// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/remoteday/rd-api-go/src/config"
	"github.com/remoteday/rd-api-go/src/db"
	"github.com/remoteday/rd-api-go/src/platform"
)

import (
	_ "github.com/lib/pq"
	_ "github.com/remoteday/rd-api-go/src/docs"
)

// Injectors from wire.go:

func InitializeApp(config2 config.AppConfig) (platform.App, error) {
	sqlDB, err := db.NewDatabaseConnection(config2)
	if err != nil {
		return platform.App{}, err
	}
	logger, err := platform.NewLogger()
	if err != nil {
		return platform.App{}, err
	}
	app := platform.NewApp(sqlDB, logger)
	return app, nil
}
