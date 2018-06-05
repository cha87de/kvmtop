FROM centos:7
RUN yum install -y libvirt-libs && yum clean all && rm -rf /var/cache/yum
ADD kvmtop /bin/kvmtop
CMD [ "/bin/kvmtop" ]