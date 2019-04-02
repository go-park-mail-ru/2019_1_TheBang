package room

import "errors"

var ErrorMaxRooms = errors.New("Max rooms limit")

var ErrorConflictName = errors.New("This name have already used")
