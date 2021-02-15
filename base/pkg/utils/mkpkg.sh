#!/bin/sh
set -eu

mkdir -vp /home/uws/build
build=$(mktemp -d /home/uws/build/pkg-XXXXXXXX)
files=/home/uws/src
dstdir=/home/uws/build

cd /uws
rm -vrf etc share include

install -v -d etc
install -v -m 0644 ${files}/etc/pkg.conf etc/

cat ${files}/manifest >${build}/+MANIFEST

exit 0
