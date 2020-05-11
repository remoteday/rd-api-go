//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/remoteday/rd-api-go/src/config"
	"github.com/remoteday/rd-api-go/src/db"
	"github.com/remoteday/rd-api-go/src/platform"
)

// InitializeApp -
func InitializeApp(config config.AppConfig) (platform.App, error) {
	wire.Build(db.NewDatabaseConnection, platform.NewApp, platform.NewLogger)
	return platform.App{}, nil
}
