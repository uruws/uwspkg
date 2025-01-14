DESTDIR ?=
PREFIX ?= /uws
BUILDDIR ?= ../../build/devel/uwspkg-build
GOPATH := /cache/go
GOCACHE := /tmp/go-cache

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

.PHONY: depends
depends:
	@cd $(BUILDDIR) && GOPATH=$(GOPATH) GOCACHE=$(GOCACHE) go mod download -x

.PHONY: build
build:
	@$(MAKE) -C $(BUILDDIR) uwspkg-build GOPATH=$(GOPATH) GOCACHE=$(GOCACHE) \
		BUILD_TAGS='-tags uwspkg_build_package'
	@echo "$(BUILDDIR)/_build/cmd/uwspkg-build created!"

.PHONY: check
check:
	@env | sort
	@doas install -v -d -m 0755 /etc/schroot/uwspkg-golang
	@$(MAKE) -C $(BUILDDIR) test GOPATH=$(GOPATH) GOCACHE=$(GOCACHE)

.PHONY: install
install:
	@$(MAKE) -C $(BUILDDIR) install-uwspkg-build DESTDIR=$(DESTDIR) PREFIX=$(PREFIX)
