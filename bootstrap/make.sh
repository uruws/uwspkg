#!/bin/sh
set -eu
PKG=${1:?'pkg version?'}
destdir=$(mktemp -d -p /tmp pkg-${PKG}-XXXXXXXX)
version=`echo -n $(cat ./VERSION)`

builddir=${PWD}/build/bootstrap-${version}

install -v -d -m 0755 ${destdir}/uws/etc
make -C ${PWD}/build/pkg-${PKG} install DESTDIR=${destdir} PREFIX=/uws

rm -rf ${builddir}
mkdir -vp ${builddir}
/uwspkg/libexec/internal/mkpkg \
	-manifest /build/base/uwspkg/manifest.yml \
	-pkg base/uwspkg \
	-builddir ${builddir}
/uwspkg/libexec/internal/mkpkg \
	-plist /build/base/uwspkg/manifest.yml \
	-pkg base/uwspkg \
	-destdir ${destdir} \
	-builddir ${builddir}

echo '/uws/sbin/uwspkg-bootstrap' >>${builddir}/pkg-plist
echo '/uws/lib/uwspkg/bootstrap/+MANIFEST' >>${builddir}/pkg-plist
echo '/uws/lib/uwspkg/bootstrap/pkg-plist' >>${builddir}/pkg-plist

install -v -d -m 0755 ${destdir}/uws/lib/uwspkg/bootstrap
install -v -m 0640 ${builddir}/+MANIFEST ${destdir}/uws/lib/uwspkg/bootstrap/
install -v -m 0640 ${builddir}/pkg-plist ${destdir}/uws/lib/uwspkg/bootstrap/

install -v -m 0755 ./bootstrap/uwspkg-bootstrap.sh ${destdir}/uws/sbin/uwspkg-bootstrap

tar -C ${destdir} -czf ./build/uwspkg-bootstrap-${version}.tgz ./
ls ./build/uwspkg-bootstrap-${version}.tgz

exit 0
