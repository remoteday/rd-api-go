package platform

import (
	"github.com/jmoiron/sqlx"
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
}

// Usecases -
type Usecases struct {
	Team team.UseCase
}

// NewApp -
func NewApp(dbConn *sqlx.DB, teamRepo team.Repository, teamUC team.UseCase) App {
	repos := Repositories{
		Team: teamRepo,
	}

	usecases := Usecases{
		Team: teamUC,
	}

	return App{
		DbConn:       dbConn,
		Repositories: repos,
		Usecases:     usecases,
	}
}
