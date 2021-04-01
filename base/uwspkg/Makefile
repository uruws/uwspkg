DESTDIR ?=
PREFIX ?= /uws
BUILDDIR ?= ../../build/base/uwspkg
CACHEDIR := /cache
PKG_CKSUM := $(PWD)/pkg.checksum

.PHONY: default
default: build

SOURCES := ../../LICENSE ../../README.md ../../VERSION ../../Makefile
SOURCES += ../../build.sh

.PHONY: fetch
fetch:
	@mkdir -p $(BUILDDIR)
	@rsync -ax --delete-before $(SOURCES) $(BUILDDIR)/
	@rsync -ax --delete-before ../../etc/pkg.conf $(BUILDDIR)/etc/
	@rsync -ax --delete-before ../../bin/uwspkg $(BUILDDIR)/bin/
	@$(MAKE) -C ../../ fetch BUILDDIR=$(BUILDDIR) CACHEDIR=$(CACHEDIR) \
		PKG_CKSUM=$(PKG_CKSUM)

.PHONY: depends
depends:
	@doas apt-get update -yy
	@cat debian-devel.install | xargs doas apt-get install --no-install-recommends -yy --purge

.PHONY: build
build:
	@$(MAKE) -C $(BUILDDIR) build PREFIX=$(PREFIX)

.PHONY: check
check:

.PHONY: install
install:
	@$(MAKE) -C $(BUILDDIR) install DESTDIR=$(DESTDIR) PREFIX=$(PREFIX)
