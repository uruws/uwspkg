#!/bin/sh
set -eu

export PATH=/uws/sbin:${PATH}

mkdir -vp /home/uws/build
build=$(mktemp -d /home/uws/build/pkg-XXXXXXXX)
files=/home/uws/src
dstfn=/home/uws/build/uwspkg-bootstrap.tgz

oldwd=${PWD}
cd /uws
rm -vrf etc share include

install -v -d etc
install -v -m 0644 ${files}/etc/pkg.conf etc/

cat ${files}/manifest >${build}/+MANIFEST

cd ${build}
pkg create -v -m . -r /uws
pkg register -d -m .

cd ${oldwd}
rm -rf ${build}

tar -C / -vczf ${dstfn} ./uws
echo "${dstfn} done!"
exit 0
