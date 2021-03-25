FROM uwspkg/base

LABEL mantainer="Jeremías Casteglione <jeremias.tincan@gmail.com>"
LABEL version="2021.03.25"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

RUN /root/bin/apt-distsetup.sh bullseye
RUN /root/bin/apt-install.sh golang-go make less

RUN mkdir -vp /go/src/uwspkg /go/.cache && chown -vR uws:uws /go
VOLUME /go/src/uwspkg

USER uws:uws
WORKDIR /home/uws

ENV USER uws
ENV HOME /home/uws

ENV GOPATH /go
ENV GOENV /go/src/uwspkg/go.env
ENV GOCACHE /go/.cache

VOLUME /go/.cache

RUN go env \
	&& go version

WORKDIR /go/src/uwspkg
CMD exec /usr/local/bin/uws-login.sh
