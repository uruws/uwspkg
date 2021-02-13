FROM alpine:3

LABEL mantainer="Jeremías Casteglione <jeremias.tincan@gmail.com>"
LABEL version="2021.02.13"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

RUN apk upgrade -U --no-cache \
	&& apk add --no-cache bash sudo make git musl-dev gcc pkgconf \
		openssl openssl-dev zlib zlib-dev linux-headers scdoc \
		lua5.3 lua5.3-dev lua5.3-lzlib

RUN printf 'uws\tALL=(ALL) NOPASSWD:ALL' >/etc/sudoers.d/uws

ARG UWS_UID=1000
ARG UWS_GID=1000

RUN addgroup -g ${UWS_GID} uws \
	&& adduser -h /home/uws -g uws -s /bin/bash -G uws -D -u ${UWS_UID} uws \
	&& chmod -v 0750 /home/uws

RUN mkdir -vp /uwspkg/src \
	&& chown -vR uws:uws /uwspkg \
	&& chmod -vR 0750 /uwspkg

VOLUME /uwspkg/src

USER uws:uws
WORKDIR /home/uws

ENV USER uws
ENV HOME /home/uws

ARG UWS_UMASK=0027

RUN printf 'umask %s\n' ${UWS_UMASK} >>/home/uws/.profile
RUN printf "export PS1='%s '\n" '\u@\h:\W\$' >>/home/uws/.profile

WORKDIR /uwspkg/src
CMD exec /bin/bash -l
