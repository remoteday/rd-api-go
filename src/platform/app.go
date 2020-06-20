package platform

import (
	"github.com/jmoiron/sqlx"
	"github.com/remoteday/rd-api-go/src/room"
	"github.com/remoteday/rd-api-go/src/team"
)

// App -
type App struct {
	DbConn       *sqlx.DB
	Repositories Repositories
	Usecases     Usecases
}

// Repositories -
type Repositories struct {
	Team team.Repository
	Room room.Repository
}

// Usecases -
type Usecases struct {
	Team team.UseCase
	Room room.UseCase
}

// NewApp -
func NewApp(dbConn *sqlx.DB, teamRepo team.Repository, teamUC team.UseCase, roomRepo room.Repository, roomUC room.UseCase) App {
	repos := Repositories{
		Team: teamRepo,
		Room: roomRepo,
	}

	usecases := Usecases{
		Team: teamUC,
		Room: roomUC,
	}

	return App{
		DbConn:       dbConn,
		Repositories: repos,
		Usecases:     usecases,
	}
}
