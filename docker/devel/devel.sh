#!/bin/sh
exec docker run -it --rm --name uwspkg-devel \
	--hostname devel.uwspkg.local \
	-v ${PWD}/build:/home/uws/build \
	-v ${PWD}:/home/uws/src/uwspkg \
	-u uws uwspkg/devel $@
