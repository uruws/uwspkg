BUILDDIR ?= ./build
CACHEDIR ?= ./build/cache
PKG_CKSUM ?= $(PWD)/base/uwspkg/pkg.checksum

PKG := 1.16.3

.PHONY: default
default: all

.PHONY: clean
clean:
	@rm -rvf ./build ./tmp

.PHONY: all
all: fetch build

.PHONY: setup
setup:
	@./devel/setup.sh

.PHONY: fetch
fetch:
	@mkdir -p $(BUILDDIR) $(CACHEDIR)
	@test -s $(CACHEDIR)/pkg-$(PKG).tgz || \
			wget -O $(CACHEDIR)/pkg-$(PKG).tgz \
				https://github.com/freebsd/pkg/archive/$(PKG).tar.gz
	@cd $(CACHEDIR) && sha256sum -c $(PKG_CKSUM)
	@tar -C $(BUILDDIR) -xzf $(CACHEDIR)/pkg-$(PKG).tgz

.PHONY: build
build:
	@BUILDDIR=$(BUILDDIR) PKG=$(PKG) ./build.sh

.PHONY: install
install:
	@$(MAKE) -C $(BUILDDIR)/pkg-$(PKG) install DESTDIR=$(DESTDIR) PREFIX=$(PREFIX)
