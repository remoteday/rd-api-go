//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/remoteday/rd-api-go/src/config"
	"github.com/remoteday/rd-api-go/src/db"
	"github.com/remoteday/rd-api-go/src/platform"
	"github.com/remoteday/rd-api-go/src/team"
)

// InitializeApp -
func InitializeApp(config config.AppConfig) (platform.App, error) {
	wire.Build(db.NewDatabaseConnection, platform.NewApp, team.NewTeamRepository, team.NewTeamUseCase)
	return platform.App{}, nil
}
