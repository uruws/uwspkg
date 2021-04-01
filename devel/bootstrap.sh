#!/bin/sh
set -eu
cd ./go
make uwspkg-build
doas ./_build/cmd/uwspkg-build -bootstrap
exit 0
