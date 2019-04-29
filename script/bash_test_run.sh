#!/bin/bash

go test -coverprofile=out.cover -coverpkg=$(go list ./...  | grep -v "test" |   tr '\n' ',') ./test/...
go tool cover -func=out.cover
go tool cover -html=out.cover