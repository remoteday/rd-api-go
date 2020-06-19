package team

import (
	"time"

	"github.com/google/uuid"
)

// DTO -
type DTO struct {
	ID        uuid.UUID `json:"id,string,omitempty"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateDTO struct {
	Name string `json:"name"`
}
