package room

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

// Repository -
type Repository struct {
	DbConn *sqlx.DB
}

// NewRoomRepository -
func NewRoomRepository(dbConn *sqlx.DB) Repository {
	return Repository{
		DbConn: dbConn,
	}
}

// FindByID -
func (r *Repository) FindByID(ctx context.Context, ID uuid.UUID) (Room, error) {
	query := "SELECT * FROM rooms WHERE id = $1"

	t := Room{}

	if err := r.DbConn.GetContext(ctx, &t, query, ID); err != nil {
		if err == sql.ErrNoRows {
			return Room{}, common.ErrNotFound
		}
		return Room{}, errors.Wrap(err, "selecting single room")
	}

	return t, nil
}

// FindAll -
func (r *Repository) FindAll(ctx context.Context) ([]Room, error) {
	query := "SELECT * FROM rooms WHERE status = 'active'"
	rooms := []Room{}
	if err := r.DbConn.SelectContext(ctx, &rooms, query); err != nil {
		return nil, errors.Wrap(err, "selecting rooms")
	}
	return rooms, nil
}

// Create -
func (r *Repository) Create(ctx context.Context, room Room) (Room, error) {
	payload := Room{
		Name:      room.Name,
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	const query = `INSERT INTO rooms
		(name, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4)  RETURNING id;`

	var lastInsertID uuid.UUID

	rows, err := r.DbConn.QueryContext(ctx, query, payload.Name, payload.Status, payload.CreatedAt, payload.UpdatedAt)

	if err != nil {
		return Room{}, errors.Wrap(err, "updating a room")
	}

	for rows.Next() {
		err = rows.Scan(&lastInsertID)
		if err != nil {
			return Room{}, errors.Wrap(err, "updating a room - no affected rooms")
		}
	}

	payload.ID = lastInsertID

	return payload, nil
}

// Update -
func (r *Repository) Update(ctx context.Context, ID uuid.UUID, room Room) (Room, error) {
	const query = `UPDATE rooms 
					SET "name" = $2, "updated_at" = $3
				   WHERE id = $1`
	room.UpdatedAt = time.Now()
	res, err := r.DbConn.ExecContext(ctx, query, room.ID, room.Name, room.UpdatedAt)

	if err != nil {
		return Room{}, err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return Room{}, err
	}

	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)

		return Room{}, err
	}

	return room, nil
}

// Delete -
func (r *Repository) Delete(ctx context.Context, ID uuid.UUID) error {
	const query = `UPDATE rooms SET "status" = 'deleted' WHERE id = $1`

	if _, err := r.DbConn.ExecContext(ctx, query, ID); err != nil {
		return errors.Wrapf(err, "deleting room %s", ID)
	}

	return nil
}
