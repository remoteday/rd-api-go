package room

import (
	"time"

	"github.com/google/uuid"
)

// Room -
type Room struct {
	ID        uuid.UUID
	Name      string
	Status    string
	TeamID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
