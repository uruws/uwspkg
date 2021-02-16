.PHONY: default
default: all

.PHONY: prune
prune:
	@docker system prune -f

.PHONY: clean
clean:
	@rm -rvf ./build ./tmp

.PHONY: all
all: docker/base docker/build docker/devel base/pkg

.PHONY: docker/base
docker/base:
	@./docker/base/build.sh

.PHONY: docker/build
docker/build:
	@./docker/build/build.sh

.PHONY: docker/devel
docker/devel:
	@./docker/devel/build.sh

.PHONY: base/pkg
base/pkg: docker/build
	@./base/pkg/build.sh
	@./base/pkg/make.sh

.PHONY: check
check: clean prune base/pkg
	@./docker/check/build.sh
	@./docker/check/run.sh
