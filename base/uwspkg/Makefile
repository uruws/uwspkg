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
	@$(MAKE) -C ../../ fetch BUILDDIR=$(BUILDDIR) CACHEDIR=$(CACHEDIR) \
		PKG_CKSUM=$(PKG_CKSUM)

.PHONY: build
build:
	@$(MAKE) -C $(BUILDDIR) build

.PHONY: check
check:

.PHONY: install
install:
