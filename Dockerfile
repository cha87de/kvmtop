FROM ubuntu:18.04
RUN apt-get update && apt-get install -y libvirt-bin
ADD kvmtop /bin/kvmtop
CMD [ "/bin/kvmtop" ]