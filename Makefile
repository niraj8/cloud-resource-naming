test:
	go test ./...

test/coverage:
	go test ./... -coverprofile=coverage.out
