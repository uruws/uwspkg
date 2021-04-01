DESTDIR ?=
PREFIX ?= /uws
BUILDDIR ?= ./build
CACHEDIR ?= ./build/cache
PKG_CKSUM ?= $(PWD)/base/uwspkg/pkg.checksum
SRV_UWSPKG ?= /srv/uwspkg

PKG := 1.16.3

.PHONY: all
all: fetch build

.PHONY: clean
clean:
	@rm -rvf ./build ./tmp

.PHONY: distclean
distclean:
	@rm -rf /etc/schroot/internal-uwspkg /etc/schroot/bootstrap-uwspkg \
		/etc/schroot/uwspkg-build-* /etc/schroot/uwspkg-* \
		/etc/schroot/chroot.d/uwspkg*.conf

.PHONY: setup-clean
setup-clean: distclean
	@rm -rf $(SRV_UWSPKG)/build/*

.PHONY: setup-distclean
setup-distclean: setup-clean
	@rm -rf $(SRV_UWSPKG)/cache/* $(SRV_UWSPKG)/chroot/*

.PHONY: bootstrap
bootstrap:
	@$(MAKE) -C go internal-mkpkg
	@./devel/bootstrap.sh $(PKG)

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
	@rm -rf $(BUILDDIR)/pkg-$(PKG)
	@tar -C $(BUILDDIR) -xzf $(CACHEDIR)/pkg-$(PKG).tgz

.PHONY: build
build:
	@BUILDDIR=$(BUILDDIR) PKG=$(PKG) PREFIX=$(PREFIX) ./build.sh

.PHONY: install
install:
	@install -v -d -m 0755 $(DESTDIR)$(PREFIX)/etc $(DESTDIR)$(PREFIX)/usr/local/bin
	@install -v -m 0644 ./etc/pkg.conf $(DESTDIR)$(PREFIX)/etc/
	@install -v -m 0755 ./bin/uwspkg $(DESTDIR)$(PREFIX)/usr/local/bin/
	@$(MAKE) -C $(BUILDDIR)/pkg-$(PKG) install DESTDIR=$(DESTDIR) PREFIX=$(PREFIX)
