package room

// ToRoom -
func ToRoom(roomDTO DTO) Room {
	return Room{
		ID:        roomDTO.ID,
		Name:      roomDTO.Name,
		Status:    roomDTO.Status,
		CreatedAt: roomDTO.CreatedAt,
		UpdatedAt: roomDTO.UpdatedAt,
	}
}

// ToRoomDTO -
func ToRoomDTO(room Room) DTO {
	return DTO{ID: room.ID, Name: room.Name, Status: room.Status, CreatedAt: room.CreatedAt, UpdatedAt: room.UpdatedAt}
}

// ToRoomDTOs -
func ToRoomDTOs(rooms []Room) []DTO {
	roomsdtos := make([]DTO, len(rooms))

	for i, itm := range rooms {
		roomsdtos[i] = ToRoomDTO(itm)
	}

	return roomsdtos
}
