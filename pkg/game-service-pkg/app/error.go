package app

import "errors"

var (
	ErrorMaxRoomsLimit = errors.New("Rooms limit")
	ErrorRoomNotFound  = errors.New("There id no this room")
)
