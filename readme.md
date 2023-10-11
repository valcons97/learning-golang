## Initial Repo for learning basic with golang

`go run learning-go/main.go`

## How to use go mod init

`go mod init github.com/valcons97/filename`

## How to use go module in another folder

`go work init`
`go work use ./path-to-module  ./path-to-module2`

make sure inside the folder there is `go.mod`
https://stackoverflow.com/a/58524450

## How to run test from root to specific package

Run test
`go test -v ./learning-go-test -run Test`

Run test coverage then show up which part is covered and not in browser
`go test ./learning-go-test -coverprofile=coverage.out && go tool cover -html=coverage.out `
for windows vs code
`go test ./learning-go-test -coverprofile=coverage.out; go tool cover -html=coverage`
