#!/bin/sh
set -eu
verfn=/home/uws/build/uwspkg-bootstrap.version
if ! test -s ${verfn}; then
	echo "${verfn}: file not found!" >&2
	exit 1
fi
export UWSPKG_VERSION=$(cat ${verfn})
check_dir=$(dirname $0)
run-parts -v ${check_dir}
exit 0
