#!/bin/sh
set -eu

# configure devel profile

prof=internal-uwspkg-devel

doas rm -rf /etc/schroot/${prof}
doas cp -va /etc/schroot/internal-uwspkg /etc/schroot/${prof}

# install bootstrap-uwspkg deps

debpkg=$(cat ./go/etc/schroot/internal-uwspkg/debian.install)
debpkg="${debpkg} $(doas cat ./go/etc/schroot/uwspkg-build/debian.install)"
debpkg="${debpkg} man less vim-tiny"

sess=$(schroot -c ${prof} -b)

cleanup() {
	schroot -c ${sess} -e
}

trap cleanup INT EXIT

schroot_sess="schroot -c ${sess} -r"

${schroot_sess} -d /root -u root -- apt-get -q update -yy
echo ${debpkg} | xargs ${schroot_sess} -d /root -u root -- \
	apt-get -q install -yy --purge --no-install-recommends

echo 'permit nopass keepenv setenv { PATH } :uws as root' |
	${schroot_sess} -d /root -u root -- tee /etc/doas.conf

${schroot_sess} -d /uwspkg/src -- /bin/bash

exit 0
