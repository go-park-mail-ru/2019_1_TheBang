#!/bin/bash

#  запускать из корня проекта!

go run cmd/auth-service/main.go . &
go run cmd/game-service/main.go . &
go run cmd/main-service/main.go . &


# [1] 81984
# [2] 81985
# [3] 81986