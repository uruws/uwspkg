#!/bin/sh
set -eu
pkgdir=${BUILDDIR}/pkg-${PKG}
cd ${pkgdir}
CC=clang ./configure --prefix=/uws
make check CC=clang CFLAGS='-D_XOPEN_SOURCE=700'
exit 0
