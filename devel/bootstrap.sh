#!/bin/sh
set -eu

PKG=${1:?'pkg version?'}

cd ./go
make uwspkg-build
doas ./_build/cmd/uwspkg-build -bootstrap
cd ../

doas rm -rf /etc/schroot/bootstrap-uwspkg
doas cp -va /etc/schroot/uwspkg-clang /etc/schroot/bootstrap-uwspkg

echo "${PWD} /build none rw,bind 0 0" | doas tee -a /etc/schroot/bootstrap-uwspkg/fstab

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

${schroot_sess} -- make PWD=/build
${schroot_sess} -- ./bootstrap/make.sh ${PKG}

exit 0
