package team

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/remoteday/rd-api-go/src/common"
)

// UseCase -
type UseCase struct {
	TeamRepository Repository
}

// NewTeamUseCase -
func NewTeamUseCase(TeamRepository Repository) UseCase {
	return UseCase{
		TeamRepository: TeamRepository,
	}
}

// FindByID -
func (u *UseCase) FindByID(ctx context.Context, ID uuid.UUID) (Team, error) {
	return u.TeamRepository.FindByID(ctx, ID)
}

// FindAll -
func (u *UseCase) FindAll(ctx context.Context) ([]Team, error) {
	return u.TeamRepository.FindAll(ctx)
}

// Create -
func (u *UseCase) Create(ctx context.Context, team Team) (Team, error) {
	return u.TeamRepository.Create(ctx, team)
}

// Update -
func (u *UseCase) Update(ctx context.Context, ID uuid.UUID, team Team) (Team, error) {
	t, err := u.FindByID(ctx, ID)
	if err != nil {
		return Team{}, common.Errors["ErrNotFound"]
	}

	t.Name = team.Name

	_, err = u.TeamRepository.Update(ctx, ID, t)
	fmt.Println("err", err, t)
	if err != nil {
		return Team{}, err
	}
	return t, nil
}

// Delete -
func (u *UseCase) Delete(ctx context.Context, ID uuid.UUID) error {
	return u.TeamRepository.Delete(ctx, ID)
}
