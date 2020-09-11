package team

import (
	"context"

	"github.com/google/uuid"
)

// Repository -
type Repository interface {
	FindByID(ctx context.Context, ID uuid.UUID) (Team, error)
	FindAll(ctx context.Context) ([]Team, error)
	Create(ctx context.Context, team Team) (Team, error)
	Update(ctx context.Context, ID uuid.UUID, team Team) (Team, error)
	Delete(ctx context.Context, ID uuid.UUID) error
}
