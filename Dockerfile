FROM alpine:latest
RUN apk update
RUN apk add libvirt-client ncurses5-libs git gettext curl bash make
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN ln -s /usr/lib/libncurses.so.5 /usr/lib/libtinfo.so.5

# use bpkg to handle complex bash entrypoints
RUN curl -Lo- "https://raw.githubusercontent.com/bpkg/bpkg/master/setup.sh" | bash
RUN bpkg install cha87de/bashutil -g

# copy entrypoint
RUN mkdir -p /opt/docker-init
ADD init /opt/docker-init

# add kvmtop binaries
ADD dist/linux_amd64/kvmtop /bin/kvmtop
ADD dist/linux_amd64/kvmprofiler /bin/kvmprofiler

# set parameters
ENV PARAMS "-c qemu:///system --printer=text --cpu --mem --net --disk"
ENV PROFILER_PARAMS "--states 4 --history 1 --filterstddevs 12 --outputFreq 20"

# start from init folder
WORKDIR /opt/docker-init
ENTRYPOINT ["./entrypoint"]