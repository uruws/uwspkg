FROM uwspkg/build

LABEL mantainer="Jeremías Casteglione <jeremias.tincan@gmail.com>"
LABEL version="2021.02.18"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

RUN /root/bin/apt-install.sh sudo \
	&& printf 'uws\tALL=(ALL) NOPASSWD:ALL' >/etc/sudoers.d/uws \
	&& chmod -v 0440 /etc/sudoers.d/uws

RUN /root/bin/apt-install.sh man less vim-tiny

USER uws:uws
WORKDIR /home/uws

ENV USER uws
ENV HOME /home/uws

RUN mkdir -vp /home/uws/src/uwspkg
VOLUME /home/uws/src/uwspkg

WORKDIR /home/uws/src/uwspkg
CMD exec /usr/local/bin/uws-login.sh
