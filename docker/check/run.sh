#!/bin/sh
mkdir -vp ${PWD}/build
exec docker run --rm --network none --name uwspkg-check \
	--hostname check.uwspkg.local \
	-v ${PWD}/build:/home/uws/build \
	-v ${PWD}/check:/home/uws/check \
	-u uws uwspkg/check $@
