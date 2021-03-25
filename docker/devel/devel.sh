#!/bin/sh
exec docker run -it --rm --name uwspkg-devel \
	--hostname devel.uwspkg.local \
	-v ${PWD}/base/uwspkg/files:/home/uws/src:ro \
	-v ${PWD}:/opt/src/uwspkg \
	-e UWSPKG_VERSION=$(cat VERSION) \
	-u uws uwspkg/devel $@
