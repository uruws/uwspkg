FROM uwspkg/base

LABEL mantainer="Jeremías Casteglione <jeremias.tincan@gmail.com>"
LABEL version="2021.03.25"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

RUN /root/bin/apt-distsetup.sh bullseye
RUN /root/bin/apt-install.sh golang-go make less

RUN mkdir -vp /go/src/uwspkg /go/.cache \
	&& chown -vR uws:uws /go
VOLUME /go/src/uwspkg

RUN mkdir -vp /usr/local/src/base /usr/local/src/devel \
	&& chown -vR uws:uws /usr/local/src
VOLUME /usr/local/src/base
VOLUME /usr/local/src/devel

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

COPY --chown=uws:uws go.mod go.sum ./
RUN go mod download -x \
	&& rm -vf go.mod go.sum

CMD exec /usr/local/bin/uws-login.sh
