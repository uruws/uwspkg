#!/bin/sh
exec docker run -it --rm --name uwspkg-uwspkg-devel \
	--hostname uwspkg-devel.uwspkg.local \
	--entrypoint /usr/local/bin/uws-login.sh \
	-e UWSPKG_VERSION=$(cat VERSION) \
	-u uws uwspkg/base:uwspkg $@
