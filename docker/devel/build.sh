#!/bin/sh
set -eu
exec docker build $@ --rm -t uwspkg/devel ./docker/devel
