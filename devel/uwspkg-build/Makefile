DESTDIR ?=
BUILDDIR ?= ../../build/devel/uwspkg-build
GOPATH := $(BUILDDIR)/_build/go
GOCACHE := $(BUILDDIR)/_build/go-cache

.PHONY: default
default: build

SOURCES := ../../go/ ../../LICENSE
EXCFILE := ../../go/.gitignore
SOURCE_EXCLUDE :=  --exclude=/.gitignore --exclude=/uwspkg.yml --exclude=Dockerfile
SOURCE_EXCLUDE += --exclude=/docker

.PHONY: fetch
fetch:
	@mkdir -p $(BUILDDIR)
	@rsync -ax --delete-before --exclude-from=$(EXCFILE) $(SOURCE_EXCLUDE) \
		$(SOURCES) $(BUILDDIR)/
	@cd $(BUILDDIR) && GOPATH=$(GOPATH) GOCACHE=$(GOCACHE) go mod download -x

.PHONY: build
build:
	@$(MAKE) -C $(BUILDDIR) uwspkg-build GOPATH=$(GOPATH) GOCACHE=$(GOCACHE)
	@echo "$(BUILDDIR)/_build/cmd/uwspkg-build created!"

.PHONY: check
check:
	@$(MAKE) -C $(BUILDDIR) test GOPATH=$(GOPATH) GOCACHE=$(GOCACHE)
