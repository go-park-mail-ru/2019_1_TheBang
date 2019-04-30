#!/bin/bash

go run cmd/auth-service/*.go . 2> logs/auth &
go run cmd/main-service/*.go . 2> logs/main &
go run cmd/chat-service/*.go . 2> logs/chat &
go run cmd/game-service/*.go . 2> logs/game &