//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/remoteday/rd-api-go/internal/config"
	"github.com/remoteday/rd-api-go/internal/db"
	"github.com/remoteday/rd-api-go/internal/platform"
	"github.com/remoteday/rd-api-go/internal/room"
	"github.com/remoteday/rd-api-go/internal/team"
)

// InitializeApp -
func InitializeApp(config config.AppConfig) (platform.App, error) {
	wire.Build(db.NewDatabaseConnection, platform.NewApp, team.NewTeamRepository, team.NewTeamUseCase, room.NewRoomRepository, room.NewRoomUseCase)
	return platform.App{}, nil
}
