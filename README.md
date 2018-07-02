[![Build Status](https://travis-ci.org/cha87de/kvmtop.svg)](https://travis-ci.org/cha87de/kvmtop)
[![GitHub release](https://img.shields.io/github/release/cha87de/kvmtop.svg)](https://github.com/cha87de/kvmtop/releases)
[![GitHub stars](https://img.shields.io/github/stars/cha87de/kvmtop.svg?style=social&label=Stars)](https://github.com/cha87de/kvmtop)
[![Docker Pulls](https://img.shields.io/docker/pulls/cha87de/kvmtop.svg)](https://hub.docker.com/r/cha87de/kvmtop/)

## What kvmtop does
kvmtop reads utilisation metrics about virtual machines running on a KVM
hypervisor from the /proc filesystem and from the libvirtd socket.

Why yet another monitoring tool for virtual machines?

The CPU measurements are read directly from the hypervisors' Linux kernel.
kvmtop takes into account the difference between utilisation inside and
outside the virtual machine, which will differ e.g. in cases of cpu over
provisioning. kvmtop also collects utilisation values of the hypervisor for
virtual machines, to measure the overhead needed to run a virtual machine.

## Installation

Download and install the [latest version of the kvmtop
build](https://github.com/cha87de/kvmtop/releases/latest). Available formats are
the binary, Deb or Rpm packages.

Installation of kvmtop 2.0 on a Debian based system (e.g. Debian or Ubuntu):

```
wget -O /tmp/kvmtop.deb https://github.com/cha87de/kvmtop/releases/download/2.0/kvmtop_2.0_linux_amd64.deb
apt-get install -y libvirt-bin  # required dependency
dpkg -i /tmp/kvmtop.deb
```

Installation on a Rpm based system (e.g. Centos 7):

```
wget -O /tmp/kvmtop.rpm https://github.com/cha87de/kvmtop/releases/download/2.0/kvmtop_2.0_linux_amd64.rpm
yum install -y libvirt  # required dependency
rpm -Uvh /tmp/kvmtop.rpm
```

## Usage

```
Usage:
  kvmtop [OPTIONS]

Monitor virtual machine experience from outside on KVM hypervisor level

Application Options:
  -v, --version     Show version
  -f, --frequency=  Frequency (in seconds) for collecting metrics (default: 1)
  -r, --runs=       Amount of collection runs (default: -1)
  -c, --connection= connection uri to libvirt daemon (default: qemu:///system)
      --procfs=     path to the proc filesystem (default: /proc)
      --verbose     Verbose output, adds more detailed fields
      --cpu         enable cpu metrics
      --mem         enable memory metrics
      --disk        enable disk metrics
      --net         enable network metrics
      --io          enable io metrics (requires root)
  -b, --batch       [DEPRECATED: use --printer=text instead] use simple output e.g. for scripts
  -p, --printer=    the output printer to use (valid printers: ncurses, text, json) (default: ncurses)
  -o, --output=     the output channel to send printer output (valid output: stdout, file, tcp) (default: stdout)
      --target=     for output 'file' the location, for 'tcp' the url to the tcp server

Help Options:
  -h, --help        Show this help message

```

Exemplary output
```
UUID                                 name          cpu_cores cpu_total cpu_steal cpu_other_total cpu_other_steal
0dbe2ae8-1ee4-4b43-bdf3-b533dfe75486 ubuntu14.04-2 2         53        0         5               1
```

With `disk_read, disk_write, net_tx, net_rx` in MB/s.

Please note: although the connection to libvirt may work remote (e.g. via ssh), kvmtop requires access to the /proc file system of the hypervisor's operating system.

### Printers and Outputs

Printers are define the representation of the monitoring data. This can be for humans in ncurses, or for further processing text (space separated) or json.

Outputs define the location where the printers send data to. Output works for text and json printers, yet not for ncurses. The output may be a file or a remote tcp server.

Example scenarios:

```
# write monitoring data to log file
kvmtop --cpu --printer=text --output=file --target=/var/log/kvmtop.log

# send mointoring data to tcp server (e.g. logstash with tcp input)
kvmtop --cpu --printer=json --output=tcp --target=127.0.0.1:12345
```


### Use Docker container

To use the kvmtop docker image, the libvirt and procfs must be known inside the container. 

Example:

```
docker run --rm \
  -v /var/run/libvirt/libvirt-sock:/var/run/libvirt/libvirt-sock \
  -v /proc/:/proc-host/ \
  cha87de/kvmtop \
  kvmtop --printer text --procfs /proc-host --cpu --mem
```

Example with logstash:

```
docker run --rm \
  -v /var/run/libvirt/libvirt-sock:/var/run/libvirt/libvirt-sock \
  -v /proc/:/proc-host/ \
  cha87de/kvmtop \
  kvmtop --printer json --output=tcp --target=127.0.0.1:12345 --procfs /proc-host --cpu --mem --disk --net --io
```
