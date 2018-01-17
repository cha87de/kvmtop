# kvmtop [![Build Status](https://travis-ci.org/cha87de/kvmtop.svg)](https://travis-ci.org/cha87de/kvmtop)

## What kvmtop does
It reads utilisation metrics about virtual machines
running on the KVM hypervisor from different sources:
 - the Linux /proc filesystem
 - the libvirtd socket

Why yet another monitoring tool for virtual machines?

The CPU measurements are read directly from the hypervisors Linux kernel. kvmtop takes
into account the difference between CPU utilisation inside and outside the virtual machine,
which will differ e.g. in cases of cpu over provisioning. kvmtop also collects 
utilisation values of the hypervisor for virtual machines, to measure the overhead needed
to run a virtual machine.

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
  -b, --batch       use simple output e.g. for scripts

Help Options:
  -h, --help        Show this help message

```

Exemplary output
```
UUID                                 name          cpu_cores cpu_total cpu_steal disk_read disk_write net_tx net_rx
0dbe2ae8-1ee4-4b43-bdf3-b533dfe75486 ubuntu14.04-2 2         53        0         0.00      0.00       0.00   0.00
```

With `disk_read, disk_write, net_tx, net_rx` in MB/s.

## Setup developer workspace or compile kvmtop

```
# create workspace
mkdir kvmtop && cd kvmtop
export GOPATH=$(pwd)
# checkout sources
go get "github.com/cha87de/kvmtop"
# build binary
go install github.com/cha87de/kvmtop
```

# Known Bugs & Issues

## No Export of Schedstat from Kernel

In CENTOS7 Kernel, the schedstats are not exported any more [1][2].
One way of installing a kernel, which exports the necessary metrics is to use the ELRepo kernel-ml:

```
# import key
rpm --import https://www.elrepo.org/RPM-GPG-KEY-elrepo.org
# install repository
rpm -Uvh http://www.elrepo.org/elrepo-release-7.0-2.el7.elrepo.noarch.rpm
# install most current kernel-ml
yum --enablerepo=elrepo-kernel install kernel-ml
```

Reboot into newly installed kernel. Then remove the old kernels!

```
yum list kernel*
yum erase kernel
```

Take care with diskless nodes. To boot into new kernel, 
the file pxelinux.cfg/default has to be changed on the storage node.


[1] https://bugzilla.redhat.com/show_bug.cgi?id=1013225
[2] https://www.centos.org/forums/viewtopic.php?f=48&t=54049
