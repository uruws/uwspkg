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
rm -vrf etc share include

install -v -d etc
install -v -m 0644 ${files}/etc/pkg.conf etc/

ldd /uws/sbin/pkg | fgrep '=>' | cut -d '>' -f 2 | cut -d ' ' -f 2 >${build}/libs
for fn in $(cat ${build}/libs | sort -u); do
	libn=`echo $(basename $fn) | sed 's/\.so.*//'`
	src=$(dirname ${fn})/${libn}
	cp -va --no-preserve=ownership ${src}*.so* /uws/lib
done

plist=${build}/pkg-plist
echo '@owner root' >${plist}
echo '@group root' >>${plist}
echo '@mode 0755' >>${plist}
echo '/uws/sbin/pkg' >>${plist}
echo '@mode' >>${plist}
(find /uws -type f && find /uws -type l) | fgrep -v /uws/sbin/pkg | sort -u >>${plist}

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

tar -C / -czf ${dstfn} ./uws
tar -tzf ${dstfn} | sort
sha256sum ${dstfn} >${dstfn}.sha256sum

cat ${dstfn}.sha256sum
exit 0
