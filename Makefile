.PHONY: default
default: base/pkg

.PHONY: prune
prune:
	@docker system prune -f

.PHONY: clean
clean:
	@rm -rvf ./build ./tmp

.PHONY: all
all: docker/base docker/build docker/check docker/devel base/pkg

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

.PHONY: base/pkg
base/pkg: docker/build
	@./base/pkg/build.sh
	@./base/pkg/make.sh

.PHONY: check
check: build/uwspkg-bootstrap.version
	@./docker/check/build.sh
	@./docker/check/run.sh

DEPS := base/pkg/Dockerfile base/pkg/make.sh base/pkg/files/manifest
DEPS += base/pkg/files/bin/uwspkg base/pkg/files/etc/pkg.conf
DEPS += base/pkg/utils/mkpkg.sh

build/uwspkg-bootstrap.version: $(DEPS)
	@$(MAKE) base/pkg
