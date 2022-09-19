.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	go fmt ./...

bench:
	go test ./... -bench=.
