#!/bin/sh
set -eu
exec docker build $@ --rm -t uwspkg/build ./docker/build
