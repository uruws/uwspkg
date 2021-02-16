#!/bin/sh
mkdir -vp ${PWD}/build
exec docker run -it --rm --network none --name uwspkg-check-devel \
	--hostname check-devel.uwspkg.local \
	--add-host check-devel.uwspkg.local:127.10.0.1 \
	-v ${PWD}/build:/home/uws/build \
	-v ${PWD}/check:/home/uws/check \
	--entrypoint /usr/local/bin/uws-login.sh \
	-u uws uwspkg/check $@
