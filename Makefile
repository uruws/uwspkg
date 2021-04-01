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
	@rm -rvf /etc/schroot/internal-uwspkg /etc/schroot/bootstrap-uwspkg \
		/etc/schroot/uwspkg-build-* /etc/schroot/uwspkg-* \
		/etc/schroot/chroot.d/uwspkg*.conf /etc/schroot/chroot.d/*-uwspkg.conf

.PHONY: bootstrap
bootstrap:
	@./devel/bootstrap.sh

.PHONY: setup
setup: bootstrap
	@./devel/setup.sh

.PHONY: setup-clean
setup-clean: distclean
	@rm -rvf $(SRV_UWSPKG)/build/*

.PHONY: setup-distclean
setup-distclean: setup-clean
	@rm -rvf $(SRV_UWSPKG)/cache/* $(SRV_UWSPKG)/chroot/*

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
