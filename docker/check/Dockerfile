FROM uwspkg/base

LABEL mantainer="Jeremías Casteglione <jeremias.tincan@gmail.com>"
LABEL version="2021.02.16-1"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

RUN /root/bin/apt-install.sh sudo \
	&& printf 'uws\tALL=(ALL) NOPASSWD:ALL' >/etc/sudoers.d/uws \
	&& chmod -v 0440 /etc/sudoers.d/uws

RUN /root/bin/apt-install.sh libbsd0 libarchive13

USER uws:uws
WORKDIR /home/uws

ENV USER uws
ENV HOME /home/uws

RUN mkdir -vp /home/uws/build /home/uws/check
VOLUME /home/uws/build
VOLUME /home/uws/check

ENTRYPOINT exec /home/uws/check/run.sh
