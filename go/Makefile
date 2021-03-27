.PHONY: default
default:

.PHONY: clean
clean:
	@rm -rfv ./_build

.PHONY: fmt
fmt:
	@gofmt -s -w -l .

.PHONY: check
check:
	@go test -i ./...
	@go test ./...

.PHONY: uwspkg-build
uwspkg-build:
	@go build -i -o ./_build/cmd/uwspkg-build ./cmd/uwspkg-build
