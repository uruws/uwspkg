#!/bin/sh
set -eu

PKG=${1:?'pkg version?'}

debpkg=$(cat ./base/uwspkg/debian-devel.install)

sess=$(schroot -c bootstrap-uwspkg -b)

cleanup() {
	schroot -c ${sess} -e
}

trap cleanup INT EXIT

schroot_sess="schroot -d /build -c ${sess} -r"

${schroot_sess} -u root -- apt-get -q update -yy
echo ${debpkg} | xargs ${schroot_sess} -u root -- \
	apt-get -q install -yy --purge --no-install-recommends

${schroot_sess} -- ./bootstrap/make.sh ${PKG}

exit 0
