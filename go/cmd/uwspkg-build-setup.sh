#!/bin/sh
set -eu
PATH=/usr/sbin:$PATH
REPO=http://deb.debian.org/debian
SRV=/srv/uwspkg
mkdir -vp ${SRV}
debinst='debootstrap --variant=minbase'

if ! test -d ${SRV}/chroot/debian-buster; then
	${debinst} buster ${SRV}/chroot/debian-buster ${REPO}
	oldwd=${PWD}
	cd ${SRV}/chroot/debian-buster
	rm -rfv ./var/lib/apt/lists/* ./var/cache/apt/archives/*.deb \
		./var/cache/apt/*cache.bin
	cd ${oldwd}
fi

mkdir -vp ${SRV}/build ${SRV}/union/overlay ${SRV}/union/underlay

schsrc=./etc/schroot
schdst=/etc/schroot
schroot_conf=${schsrc}/chroot.d/uwspkg.conf
PROFILES='default build'

if test -s ${schroot_conf}; then
	install -v -C -m 0644 ${schroot_conf} ${schdst}/chroot.d/uwspkg.conf
	for prof in ${PROFILES}; do
		dst="${schdst}/uwspkg-${prof}"
		rm -rf ${dst}
		install -m 0755 -d ${dst}
		install -m 0644 ${schsrc}/uwspkg-${prof}/* ${dst}
	done
fi

exit 0
