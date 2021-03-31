BUILDDIR ?= ./build
CACHEDIR ?= ./build/cache
PKG_CKSUM ?= $(PWD)/base/uwspkg/pkg.checksum

PKG := 1.16.3

.PHONY: default
default: base/uwspkg

.PHONY: prune
prune:
	@docker system prune -f

.PHONY: clean
clean:
	@rm -rvf ./build ./tmp

.PHONY: all
all: docker/base docker/build docker/check docker/devel base/uwspkg go/docker

.PHONY: docker/base
docker/base:
	@./docker/base/build.sh

.PHONY: docker/build
docker/build:
	@./docker/build/build.sh

.PHONY: docker/check
docker/check:
	@./docker/check/build.sh

.PHONY: docker/devel
docker/devel:
	@./docker/devel/build.sh

.PHONY: base/uwspkg
base/uwspkg: docker/build
	@./base/uwspkg/build.sh
	@./base/uwspkg/make.sh

.PHONY: check
check: go/check build/uwspkg-bootstrap.version
	@./docker/check/build.sh
	@./docker/check/run.sh

DEPS := base/uwspkg/Dockerfile base/uwspkg/make.sh base/uwspkg/files/manifest
DEPS += base/uwspkg/files/bin/uwspkg base/uwspkg/files/etc/pkg.conf
DEPS += base/uwspkg/utils/mkpkg.sh

build/uwspkg-bootstrap.version: $(DEPS)
	@$(MAKE) base/uwspkg

.PHONY: go/docker
go/docker:
	@./go/docker/build.sh

.PHONY: go/check
go/check:
	@./go/docker/check.sh

.PHONY: fetch
fetch:
	@mkdir -p $(BUILDDIR) $(CACHEDIR)
	@test -s $(CACHEDIR)/pkg-${PKG}.tgz || \
			wget -O $(CACHEDIR)/pkg-${PKG}.tgz \
				https://github.com/freebsd/pkg/archive/${PKG}.tar.gz
	@cd $(CACHEDIR) && sha256sum -c $(PKG_CKSUM)
	@tar -C $(BUILDDIR) -xzf $(CACHEDIR)/pkg-${PKG}.tgz
