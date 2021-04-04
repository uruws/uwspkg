DESTDIR ?=
PREFIX ?= /uws

PKGNAME := perms
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
	@echo 1 >$(BUILDDIR)/test/$(PKGNAME)-1.pkg
	@echo 1 >$(BUILDDIR)/test/$(PKGNAME)-2.pkg

.PHONY: check
check:
	@stat $(BUILDDIR)/test/$(PKGNAME).pkg
	@stat $(BUILDDIR)/test/$(PKGNAME)-1.pkg
	@stat $(BUILDDIR)/test/$(PKGNAME)-2.pkg

.PHONY: install
install:
	@install -d -v -m 0755 $(DESTDIR)$(PREFIX)/test/$(PKGNAME).dir
	@install -d -v -m 0755 $(DESTDIR)$(PREFIX)/test/$(PKGNAME)-1.dir
	@install -v -m 0644 $(BUILDDIR)/test/$(PKGNAME)*.pkg $(DESTDIR)$(PREFIX)/test/
