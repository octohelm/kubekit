gen:
	go run ./tool/internal/cmd/gen

test:
	go test -v -failfast ./...