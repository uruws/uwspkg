.PHONY: default
default: all

.PHONY: prune
prune:
	@docker system prune -f

.PHONY: all
all: docker/base docker/build docker/devel

.PHONY: docker/base
docker/base:
	@./docker/base/build.sh

.PHONY: docker/build
docker/build:
	@./docker/build/build.sh

.PHONY: docker/devel
docker/devel:
	@./docker/devel/build.sh
