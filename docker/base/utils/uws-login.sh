#!/bin/sh
export USER=uws
export HOME=/home/uws
if test -d /uws/sbin; then
	export PATH=/uws/sbin:${PATH}
fi
exec /bin/bash -l
