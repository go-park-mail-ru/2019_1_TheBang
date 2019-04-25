#!/bin/bash

go run cmd/auth-service/*.go . &
go run cmd/game-service/*.go . &
go run cmd/main-service/*.go . &