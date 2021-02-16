#!/bin/sh
exec docker run -it --rm --name uwspkg-pkg-devel \
	--hostname pkg-devel.uwspkg.local \
	--entrypoint /usr/local/bin/uws-login.sh \
	-u uws uwspkg/pkg $@
