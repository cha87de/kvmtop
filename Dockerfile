FROM alpine:latest
RUN apk update
RUN apk add libvirt-client ncurses5-libs
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN ln -s /usr/lib/libncurses.so.5 /usr/lib/libtinfo.so.5
ADD kvmtop /bin/kvmtop
CMD [ "/bin/kvmtop" ]