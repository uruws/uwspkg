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
	@install -d -m 0755 $(DESTDIR)/etc
	@install -d -m 0755 $(DESTDIR)/etc/schroot
	@install -d -m 0755 $(DESTDIR)/etc/schroot/chroot.d
	@install -v -m 0644 ./etc/schroot/chroot.d/*.conf \
		$(DESTDIR)/etc/schroot/chroot.d/
	@install -d -m 0755 $(DESTDIR)/etc/schroot/internal-uwspkg
	@install -v -m 0644 ./etc/schroot/internal-uwspkg/* \
		$(DESTDIR)/etc/schroot/internal-uwspkg/
	@for profd in `ls -d etc/schroot/uwspkg-*`; do \
		install -d -m 0755 $(DESTDIR)/$${profd}; \
		install -v -m 0644 ./$${profd}/* $(DESTDIR)/$${profd}/; \
	done
	@install -d -m 0755 $(DESTDIR)$(PREFIX)/libexec
	@install -d -m 0755 $(DESTDIR)$(PREFIX)/libexec/uwspkg
	@for libxd in `ls libexec/utils`; do \
		install -d -m 0755 $(DESTDIR)$(PREFIX)/libexec/uwspkg/$${libxd}; \
		install -v -m 0755 ./libexec/utils/$${libxd}/* \
			$(DESTDIR)$(PREFIX)/libexec/uwspkg/$${libxd}/; \
	done
	@chmod 0644 $(DESTDIR)$(PREFIX)/libexec/uwspkg/internal/env.export
