#!/bin/sh
set -eu
bootstrap_tgz=${PWD}/build/uwspkg-bootstrap.tgz
test -s ${bootstrap_tgz} || {
	echo "${bootstrap_tgz}: file not found" >&2
	exit 1
}
cd ./go
make uwspkg-build
doas ./_build/cmd/uwspkg-build -setup
exit 0
