#!/bin/sh
set -eu
version=`echo -n $(cat ./VERSION)`
bootstrap_tgz=${PWD}/build/uwspkg-bootstrap-${version}.tgz
test -s ${bootstrap_tgz} || {
	echo "${bootstrap_tgz}: file not found" >&2
	exit 1
}
cd ./go
make uwspkg-build
doas ./_build/cmd/uwspkg-build -setup
exit 0
