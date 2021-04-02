DESTDIR ?=
PREFIX ?= /uws
BUILDDIR := ../../build/test/base

.PHONY: default
default: build

.PHONY: fetch
fetch:

.PHONY: depends
depends:

.PHONY: build
build:
	@mkdir -vp $(BUILDDIR)/test
	@echo 1 >$(BUILDDIR)/test/base.pkg

.PHONY: check
check:
	@stat $(BUILDDIR)/test/base.pkg

.PHONY: install
install:
	@install -d -v -m 0755 $(DESTDIR)$(PREFIX)/test
	@install -v -m 0644 $(BUILDDIR)/test/base.pkg $(DESTDIR)$(PREFIX)/test/base.pkg
