.PHONY: default
default:

.PHONY: clean
clean:
	@rm -rfv ./_build

.PHONY: fmt
fmt:
	@gofmt -s -w -l .

.PHONY: test
test:
	@go test ./...

.PHONY: check
check:
	@go test -i ./...
	@go test ./...
	@go test -race ./...

.PHONY: uwspkg-build
uwspkg-build:
	@go build -i -o ./_build/cmd/uwspkg-build ./cmd/uwspkg-build
