FROM ubuntu:18.04
USER root

# golang general
RUN apt-get update && apt-get install -y golang-go
VOLUME /opt/gopath
WORKDIR /opt/gopath
ENV GOPATH /opt/gopath

# kvmtop specific
RUN apt-get install -y libvirt-dev pkg-config libncurses5-dev
CMD go build github.com/cha87de/kvmtop/cmd/kvmtop