package team

import (
	context "context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCanGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockRepository(ctrl)

	bg := context.Background()
	id := uuid.New()
	expected := Team{ID: id, Name: "test", Status: "active"}

	repo.EXPECT().FindByID(bg, id).Return(expected, nil)

	uc := NewTeamUseCase(repo)
	team, err := uc.FindByID(bg, id)

	assert.NoError(t, err)
	assert.Equal(t, expected, team)
}
