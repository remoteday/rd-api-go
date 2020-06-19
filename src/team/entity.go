package team

import (
	"time"

	"github.com/google/uuid"
)

// Team -
type Team struct {
	ID        uuid.UUID
	Name      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
