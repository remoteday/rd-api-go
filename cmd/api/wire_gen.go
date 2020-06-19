// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/remoteday/rd-api-go/src/config"
	"github.com/remoteday/rd-api-go/src/db"
	"github.com/remoteday/rd-api-go/src/platform"
	"github.com/remoteday/rd-api-go/src/team"
)

import (
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Injectors from wire.go:

func InitializeApp(config2 config.AppConfig) (platform.App, error) {
	sqlxDB, err := db.NewDatabaseConnection(config2)
	if err != nil {
		return platform.App{}, err
	}
	repository := team.NewTeamRepository(sqlxDB)
	useCase := team.NewTeamUseCase(repository)
	app := platform.NewApp(sqlxDB, repository, useCase)
	return app, nil
}
