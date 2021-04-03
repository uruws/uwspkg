DESTDIR ?=
PREFIX ?= /uws

PKGNAME := service
BUILDDIR := ../../build/test/$(PKGNAME)

.PHONY: default
default: build

.PHONY: fetch
fetch:

.PHONY: depends
depends:

.PHONY: build
build:
	@mkdir -vp $(BUILDDIR)/test
	@echo 1 >$(BUILDDIR)/test/$(PKGNAME).pkg

.PHONY: check
check:
	@stat $(BUILDDIR)/test/$(PKGNAME).pkg

.PHONY: install
install:
	@install -d -v -m 0755 $(DESTDIR)$(PREFIX)/test/$(PKGNAME).dir
	@install -v -m 0644 $(BUILDDIR)/test/$(PKGNAME).pkg $(DESTDIR)$(PREFIX)/test/$(PKGNAME).pkg
