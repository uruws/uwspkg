#!/bin/sh
set -eu
prefix=${PREFIX}
pkgdir=${BUILDDIR}/pkg-${PKG}
cd ${pkgdir}
CC=clang ./configure --prefix=${PREFIX}
make check PREFIX=${prefix} CC=clang CFLAGS='-D_XOPEN_SOURCE=700'
exit 0
