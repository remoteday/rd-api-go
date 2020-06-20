package room

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/remoteday/rd-api-go/src/common"
)

// UseCase -
type UseCase struct {
	RoomRepository Repository
}

// NewRoomUseCase -
func NewRoomUseCase(roomRepository Repository) UseCase {
	return UseCase{
		RoomRepository: roomRepository,
	}
}

// FindByID -
func (u *UseCase) FindByID(ctx context.Context, ID uuid.UUID) (Room, error) {
	return u.RoomRepository.FindByID(ctx, ID)
}

// FindAll -
func (u *UseCase) FindAll(ctx context.Context) ([]Room, error) {
	return u.RoomRepository.FindAll(ctx)
}

// Create -
func (u *UseCase) Create(ctx context.Context, room Room) (Room, error) {
	return u.RoomRepository.Create(ctx, room)
}

// Update -
func (u *UseCase) Update(ctx context.Context, ID uuid.UUID, room Room) (Room, error) {
	t, err := u.FindByID(ctx, ID)
	if err != nil {
		return Room{}, common.ErrNotFound
	}

	t.Name = room.Name

	_, err = u.RoomRepository.Update(ctx, ID, t)
	fmt.Println("err", err, t)
	if err != nil {
		return Room{}, err
	}
	return t, nil
}

// Delete -
func (u *UseCase) Delete(ctx context.Context, ID uuid.UUID) error {
	return u.RoomRepository.Delete(ctx, ID)
}
