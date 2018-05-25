FROM centos:7
RUN yum install -y libvirt-libs
ADD kvmtop /bin/kvmtop
CMD [ "/bin/kvmtop" ]