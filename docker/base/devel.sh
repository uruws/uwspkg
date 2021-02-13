#!/bin/sh
exec docker run -it --rm --name uwspkg-devel \
	--hostname devel.uwspkg.local \
	--add-host devel.uwspkg.local:127.10.0.1 \
	-v ${PWD}:/uwspkg/src \
	-u uws uwspkg/base $@
