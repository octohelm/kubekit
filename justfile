gen:
	go generate ./...

fmt:
	go tool gofumpt -w -l .

test:
	go test -v -failfast ./...