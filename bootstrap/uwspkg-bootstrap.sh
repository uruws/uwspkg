#!/bin/sh
set -eu
export PATH=/uws/sbin:${PATH}
pkgdir=/uws/lib/uwspkg/bootstrap
cd ${pkgdir}
pkg -N || pkg register -d -m . -f pkg-plist
exit 0
