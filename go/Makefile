DESTDIR ?=
PREFIX ?= /usr/local

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

.PHONY: install-uwspkg-build
install-uwspkg-build:
	@install -d -m 0755 $(DESTDIR)$(PREFIX)
	@install -d -m 0755 $(DESTDIR)$(PREFIX)/bin
	@install -v -m 0755 ./_build/cmd/uwspkg-build $(DESTDIR)$(PREFIX)/bin/
