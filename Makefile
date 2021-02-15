.PHONY: default
default: all

.PHONY: prune
prune:
	@docker system prune -f

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
base/pkg:
	@./base/pkg/build.sh
