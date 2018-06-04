# kvmtop [![Build Status](https://travis-ci.org/cha87de/kvmtop.svg)](https://travis-ci.org/cha87de/kvmtop)

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
      --cpu         enable cpu metrics
      --mem         enable memory metrics
      --disk        enable disk metrics
      --net         enable network metrics
  -b, --batch       [DEPRECATED: use --printer=text instead] use simple output
                    e.g. for scripts
  -p, --printer=    the output printer to use (valid printers: ncurses, text,
                    json) (default: ncurses)

Help Options:
  -h, --help        Show this help message

```

Exemplary output
```
UUID                                 name          cpu_cores cpu_total cpu_steal disk_read disk_write net_tx net_rx
0dbe2ae8-1ee4-4b43-bdf3-b533dfe75486 ubuntu14.04-2 2         53        0         0.00      0.00       0.00   0.00
```

With `disk_read, disk_write, net_tx, net_rx` in MB/s.

Please note: although the connection to libvirt may work remote (e.g. via ssh), kvmtop requires access to the /proc file system of the hypervisor's operating system.