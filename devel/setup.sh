#!/bin/sh
set -eu
cd ./go
make uwspkg-build
doas ./_build/cmd/uwspkg-build -setup
exit 0
