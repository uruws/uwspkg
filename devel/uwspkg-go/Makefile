DESTDIR ?=
BUILDDIR ?= ./files

.PHONY: default
default:

SOURCES := ../../go/ ../../LICENSE ../../README.md ../../VERSION
EXCFILE := ../../go/.gitignore

.PHONY: fetch
fetch:
	@rsync -vax --delete-before --exclude-from=$(EXCFILE) $(SOURCES) $(BUILDDIR)/
