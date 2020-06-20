package room

import (
	"time"

	"github.com/google/uuid"
)

// Room -
type Room struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Status    string    `db:"status"`
	TeamID    string    `db:"team_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
