#!/bin/sh
set -eu

export PATH=/uws/sbin:${PATH}

mkdir -vp /home/uws/build
build=$(mktemp -d /home/uws/build/pkg-XXXXXXXX)
files=/home/uws/src
dstfn=/home/uws/build/uwspkg-bootstrap-${UWSPKG_VERSION}.tgz
verfn=/home/uws/build/uwspkg-bootstrap.version

echo ${UWSPKG_VERSION} >${verfn}

oldwd=${PWD}

cd /uws
rm -vrf etc share include lib/pkgconfig

install -v -d etc
install -v -m 0644 ${files}/etc/pkg.conf etc/
install -v -m 0755 ${files}/bin/uwspkg /usr/local/bin/uwspkg

ldd /uws/sbin/pkg >${build}/libs
LIBS='libbsd.so libarchive.so libmd.so libxml2.so libicuuc.so libicudata.so'
for fn in ${LIBS}; do
	cp -va --no-preserve=ownership /usr/lib/x86_64-linux-gnu/${fn}* /uws/lib
done

plist=${build}/pkg-plist
echo '@owner root' >${plist}
echo '@group root' >>${plist}
echo '@mode 0755' >>${plist}
echo '/usr/local/bin/uwspkg' >>${plist}
echo '/uws/sbin/pkg' >>${plist}
echo '@mode' >>${plist}

(find /uws -type f && find /uws -type l) |
	fgrep -v /uws/sbin/pkg | fgrep -v /usr/local/bin/uwspkg | sort -u >>${plist}

manifest=${build}/+MANIFEST
cat ${files}/manifest >${manifest}
echo "version: ${UWSPKG_VERSION}" >>${manifest}

mkdir -vp /uws/var/db/pkg
echo '@dir /uws/var' >>${plist}
echo '@dir /uws/var/db' >>${plist}
echo '@dir /uws/var/db/pkg' >>${plist}

cd ${build}
cat ${manifest}
cat ${plist}

pkg create -v -o /home/uws/build -m . -p pkg-plist -r /
fakeroot pkg register -d -m .

cd ${oldwd}
rm -vf /uws/var/db/pkg/local.sqlite

fakeroot tar -C / -czf ${dstfn} ./usr/local/bin/uwspkg ./uws
tar -tzf ${dstfn} | sort
sha256sum ${dstfn} >${dstfn}.sha256sum

cat ${dstfn}.sha256sum
exit 0
