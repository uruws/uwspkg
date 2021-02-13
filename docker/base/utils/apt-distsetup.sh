#!/bin/sh
set -eu

DIST=${1:?'dist name?'}
SECT=${2:-'main contrib non-free'}

DEBURI='http://deb.debian.org/debian/'
SECURI='http://security.debian.org/debian-security'
SLF="/etc/apt/sources.list.d/${DIST}.list"

echo "deb ${DEBURI} ${DIST} ${SECT}" >${SLF}
echo "deb ${DEBURI} ${DIST}-updates ${SECT}" >>${SLF}

if test "${DIST}" = 'stable'; then
	echo "deb ${SECURI} ${DIST}/updates ${SECT}" >>${SLF}
fi

export DEBIAN_FRONTEND=noninteractive

apt-get clean
apt-get update
apt-get dist-upgrade -yy --purge

apt-get clean
apt-get autoremove -yy --purge
rm -rf /var/lib/apt/lists/* \
	/var/cache/apt/archives/*.deb \
	/var/cache/apt/*cache.bin

exit 0
