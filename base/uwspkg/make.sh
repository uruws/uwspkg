#!/bin/sh
mkdir -vp ${PWD}/build
exec docker run -it --rm --name uwspkg-make \
	--hostname make.uwspkg.local \
	-v ${PWD}/build:/home/uws/build \
	-e UWSPKG_VERSION=$(cat VERSION) \
	-u uws uwspkg/base:uwspkg
