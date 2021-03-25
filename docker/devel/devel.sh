#!/bin/sh
exec docker run -it --rm --name uwspkg-devel \
	--hostname devel.uwspkg.local \
	-v ${PWD}/build:/home/uws/build \
	-v ${PWD}/base/uwspkg/files:/home/uws/src:ro \
	-v ${PWD}:/uws/src/uwspkg:ro \
	-e UWSPKG_VERSION=$(cat VERSION) \
	-u uws uwspkg/devel $@
