#!/bin/sh
mkdir -vp ${PWD}/build
exec docker run -it --rm --network none --name uwspkg-check-devel \
	--hostname check-devel.uwspkg.local \
	-v ${PWD}/build:/home/uws/build \
	-v ${PWD}/check:/home/uws/check \
	-u uws uwspkg/check $@
