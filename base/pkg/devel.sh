#!/bin/sh
exec docker run -it --rm --name uwspkg-pkg-devel \
	--hostname pkg-devel.uwspkg.local \
	--entrypoint /usr/local/bin/uws-login.sh \
	-e UWSPKG_VERSION=$(cat VERSION) \
	-u uws uwspkg/pkg $@
