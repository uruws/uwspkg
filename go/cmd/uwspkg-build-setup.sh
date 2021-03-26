#!/bin/sh
set -eu
PATH=/usr/sbin:$PATH
REPO=http://deb.debian.org/debian
SECREPO=http://security.debian.org/debian-security
SRV=/srv/uwspkg
mkdir -vp ${SRV} ${SRV}/cache/debootstrap
debinst="debootstrap --variant=minbase --cache-dir=${SRV}/cache/debootstrap"
debinst="${debinst} --force-check-gpg"

baseroot=${SRV}/chroot/debian-buster
schroot_default='schroot -c source:uwspkg-default -d /root'

if ! test -d ${baseroot}; then
	${debinst} buster ${baseroot} ${REPO}
	oldwd=${PWD}
	cd ${baseroot}
	printf 'deb %s/ buster main\n' "${REPO}" >./etc/apt/sources.list
	printf 'deb %s/ buster-updates main\n' "${REPO}" >>./etc/apt/sources.list
	printf 'deb %s buster/updates main\n' "${SECREPO}" >>./etc/apt/sources.list
	${schroot_default} -- apt-get update -yy
	${schroot_default} -- apt-get install -yy --no-install-recommends bash
	rm -rf ./var/lib/apt/lists/* ./var/cache/apt/archives/*.deb \
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
		dstroot=`dirname ${baseroot}`/${prof}
		if test 'Xdefault' != "X${prof}"; then
			rsync -vax --delete-before ${baseroot}/ ${dstroot}/
			echo "uwspkg-${prof}" >${dstroot}/etc/debian_chroot
		fi
	done
fi

exit 0
