package team

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Repository -
type Repository struct {
	DbConn *sqlx.DB
}

// NewTeamRepository -
func NewTeamRepository(dbConn *sqlx.DB) Repository {
	return Repository{
		DbConn: dbConn,
	}
}

func (r *Repository) fetch(ctx context.Context, query string, args ...interface{}) ([]Team, error) {
	rows, err := r.DbConn.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]Team, 0)
	for rows.Next() {
		t := Team{}
		err = rows.Scan(&t.ID, &t.Name, &t.Status, &t.CreatedAt, &t.UpdatedAt)

		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

// FindByID -
func (r *Repository) FindByID(ctx context.Context, ID uuid.UUID) (Team, error) {
	query := "SELECT * FROM teams WHERE id = $1"

	row := r.DbConn.QueryRowContext(ctx, query, ID)

	t := Team{}
	err := row.Scan(&t.ID, &t.Name, &t.Status, &t.CreatedAt, &t.UpdatedAt)

	if err != nil {
		return Team{}, err
	}
	return t, nil
}

// FindAll -
func (r *Repository) FindAll(ctx context.Context) ([]Team, error) {
	query := "SELECT * FROM teams WHERE status = 'active'"
	res, err := r.fetch(ctx, query)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create -
func (r *Repository) Create(ctx context.Context, team Team) (Team, error) {
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
		return Team{}, err
	}

	for rows.Next() {
		err = rows.Scan(&lastInsertID)
		if err != nil {
			return Team{}, err
		}
	}

	payload.ID = lastInsertID

	return payload, nil
}

// Update -
func (r *Repository) Update(ctx context.Context, ID uuid.UUID, team Team) (Team, error) {
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
func (r *Repository) Delete(ctx context.Context, ID uuid.UUID) error {
	const query = `UPDATE teams SET "status" = 'deleted' WHERE id = $1`

	_, err := r.DbConn.ExecContext(ctx, query, ID)

	return err
}
