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
echo '@mode 0644' >>${plist}
(find /uws -type f && find /uws -type l) | fgrep -v /uws/sbin/pkg >>${plist}

manifest=${build}/+MANIFEST
cat ${files}/manifest >${manifest}

cd ${build}
cat ${manifest}
cat ${plist}

pkg create -v -o /home/uws/build -m . -p pkg-plist -r /
pkg register -d -m .

cd ${oldwd}

tar -C / -vczf ${dstfn} ./uws
echo "${dstfn} done!"

exit 0
