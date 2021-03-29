DESTDIR ?=
BUILDDIR ?= ../../build/devel/uwspkg-build
GOPATH := $(BUILDDIR)/_build/go
GOCACHE := $(BUILDDIR)/_build/go-cache

.PHONY: default
default:

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
