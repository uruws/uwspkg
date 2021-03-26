#!/bin/sh
set -eu
PATH=/usr/sbin:$PATH
REPO=http://deb.debian.org/debian
SRV=/srv/uwspkg
mkdir -vp ${SRV}
debinst='debootstrap --variant=minbase'
if ! test -d ${SRV}/chroot/debian-buster; then
	${debinst} buster ${SRV}/chroot/debian-buster ${REPO}
fi
mkdir -vp ${SRV}/union/overlay ${SRV}/union/underlay
schsrc=./etc/schroot
schdst=/etc/schroot
schroot_conf=${schsrc}/chroot.d/uwspkg.conf
PROFILES='default'
if test -s ${schroot_conf}; then
	install -v -C -m 0644 ${schroot_conf} ${schdst}/chroot.d/uwspkg.conf
	for prof in ${PROFILES}; do
		mkdir -vp ${schdst}/uwspkg-${prof}
		install -v -C -m 0644 ${schsrc}/uwspkg-${prof}/* ${schdst}/uwspkg-${prof}
	done
fi
exit 0
