DESTDIR ?=
BUILDDIR ?= ../../build/devel/uwspkg-build

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
