DESTDIR ?=
PREFIX ?= /uws

PKGNAME := pkg-trigger
BUILDDIR := ../../build/test/$(PKGNAME)

.PHONY: default
default: build

.PHONY: fetch
fetch:
	@cp -va ./trigger0.ucl $(BUILDDIR)/

.PHONY: depends
depends:

.PHONY: build
build:
	@mkdir -vp $(BUILDDIR)/test
	@echo 1 >$(BUILDDIR)/test/$(PKGNAME).pkg
	@echo 1 >$(BUILDDIR)/test/trigger.test

.PHONY: check
check:
	@stat $(BUILDDIR)/test/$(PKGNAME).pkg

.PHONY: install
install:
	@install -d -v -m 0755 $(DESTDIR)$(PREFIX)/test/$(PKGNAME).dir
	@install -v -m 0644 $(BUILDDIR)/test/$(PKGNAME).pkg $(DESTDIR)$(PREFIX)/test/$(PKGNAME).pkg
	@install -d -v -m 0755 $(DESTDIR)$(PREFIX)/share/pkg/trigger
	@install -v -m 0644 $(BUILDDIR)/trigger0.ucl $(DESTDIR)$(PREFIX)/share/pkg/trigger/
	@install -d -v -m 0755 $(DESTDIR)$(PREFIX)/var/pkg/trigger
	@install -v -m 0644 $(BUILDDIR)/test/trigger.test $(DESTDIR)$(PREFIX)/var/pkg/trigger/test
