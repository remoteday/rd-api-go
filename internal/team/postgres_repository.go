package team

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/remoteday/rd-api-go/internal/common"
)

// PostgresRepository -
type PostgresRepository struct {
	DbConn *sqlx.DB
}

// NewTeamRepository -
func NewTeamRepository(dbConn *sqlx.DB) Repository {
	return PostgresRepository{
		DbConn: dbConn,
	}
}

// FindByID -
func (r PostgresRepository) FindByID(ctx context.Context, ID uuid.UUID) (Team, error) {
	query := "SELECT * FROM teams WHERE id = $1"

	t := Team{}

	if err := r.DbConn.GetContext(ctx, &t, query, ID); err != nil {
		if err == sql.ErrNoRows {
			return Team{}, common.ErrNotFound
		}
		return Team{}, errors.Wrap(err, "selecting single team")
	}

	return t, nil
}

// FindAll -
func (r PostgresRepository) FindAll(ctx context.Context) ([]Team, error) {
	query := "SELECT * FROM teams WHERE status = 'active'"
	teams := []Team{}
	if err := r.DbConn.SelectContext(ctx, &teams, query); err != nil {
		return nil, errors.Wrap(err, "selecting teams")
	}
	return teams, nil
}

// Create -
func (r PostgresRepository) Create(ctx context.Context, team Team) (Team, error) {
	payload := Team{
		Name:      team.Name,
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	const query = `INSERT INTO teams
		(name, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4)  RETURNING id;`

	var lastInsertID uuid.UUID

	rows, err := r.DbConn.QueryContext(ctx, query, payload.Name, payload.Status, payload.CreatedAt, payload.UpdatedAt)

	if err != nil {
		return Team{}, errors.Wrap(err, "updating a team")
	}

	for rows.Next() {
		err = rows.Scan(&lastInsertID)
		if err != nil {
			return Team{}, errors.Wrap(err, "updating a team - no affected teams")
		}
	}

	payload.ID = lastInsertID

	return payload, nil
}

// Update -
func (r PostgresRepository) Update(ctx context.Context, ID uuid.UUID, team Team) (Team, error) {
	const query = `UPDATE teams 
					SET "name" = $2, "updated_at" = $3
				   WHERE id = $1`
	team.UpdatedAt = time.Now()
	res, err := r.DbConn.ExecContext(ctx, query, team.ID, team.Name, team.UpdatedAt)

	if err != nil {
		return Team{}, err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return Team{}, err
	}

	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)

		return Team{}, err
	}

	return team, nil
}

// Delete -
func (r PostgresRepository) Delete(ctx context.Context, ID uuid.UUID) error {
	const query = `UPDATE teams SET "status" = 'deleted' WHERE id = $1`

	if _, err := r.DbConn.ExecContext(ctx, query, ID); err != nil {
		return errors.Wrapf(err, "deleting team %s", ID)
	}

	return nil
}
