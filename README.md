[![Codacy Badge](https://api.codacy.com/project/badge/Grade/2111f329884b41efa85e4556311b1f1b)](https://app.codacy.com/app/cha87de/kvmtop?utm_source=github.com&utm_medium=referral&utm_content=cha87de/kvmtop&utm_campaign=Badge_Grade_Dashboard)
[![Build Status](https://travis-ci.org/cha87de/kvmtop.svg?branch=master)](https://travis-ci.org/cha87de/kvmtop)
[![GitHub release](https://img.shields.io/github/release/cha87de/kvmtop.svg)](https://github.com/cha87de/kvmtop/releases)
[![GitHub stars](https://img.shields.io/github/stars/cha87de/kvmtop.svg?style=social&label=Stars)](https://github.com/cha87de/kvmtop)
[![Go Report Card](https://goreportcard.com/badge/github.com/cha87de/kvmtop)](https://goreportcard.com/report/github.com/cha87de/kvmtop)
[![GoDoc](https://godoc.org/github.com/cha87de/kvmtop?status.svg)](https://godoc.org/github.com/cha87de/kvmtop)
[![Docker Pulls](https://img.shields.io/docker/pulls/cha87de/kvmtop.svg)](https://hub.docker.com/r/cha87de/kvmtop/)

## What kvmtop does
kvmtop reads utilisation metrics about virtual machines running on a KVM hypervisor from the Linux proc filesystem and from libvirt.

*Why yet another monitoring tool for virtual machines?*

Kvmtop takes into account the difference between utilisation inside and
outside the virtual machine, which differs in cases of overprovisioning. Kvmtop collects utilisation values of the hypervisor for virtual machines, to measure the overhead needed to run a virtual machine. Kvmtop will help to identify resource shortcomings, leading 
to the "noisy neighbour" effect.

The conceptual idea behind kvmtop is scientifically published and described in "Reviewing Cloud Monitoring: Towards Cloud Resource Profiling."

```
@inproceedings{hauser2018reviewing,
  title={Reviewing Cloud Monitoring: Towards Cloud Resource Profiling},
  author={Hauser, Christopher B and Wesner, Stefan},
  booktitle={2018 IEEE 11th International Conference on Cloud Computing (CLOUD)},
  pages={678--685},
  year={2018},
  organization={IEEE}
}
```

*What does kvmtop offer?*

The command line tool can be used by sysadmins, using a console ui. Text or JSON output further allows to process the monitoring data. A build in TCP output allows to send the data directly to a monitoring data sink, e.g. logstash.

## Installation

Download and install the [latest version of the kvmtop
build](https://github.com/cha87de/kvmtop/releases/latest). Available formats are the binary, Deb or Rpm packages, and a Docker image.

On Ubuntu, download the deb file and install it with `dpkg -i kvmtop_VERSION_linux_amd64.deb`. Similarly, on a rpm based system (e.g. Centos 7) with `rpm -Uvh kvmtop_VERSION_linux_amd64.rpm`.


### Docker Usage

To use the Docker image, create a container from `cha87de/kvmtop:master` for the current master build or `cha87de/kvmtop:latest` to use the latest stable release. 

```
docker run --rm \
  -v /var/run/libvirt/libvirt-sock:/var/run/libvirt/libvirt-sock \
  --privileged --pid="host" \
  cha87de/kvmtop:latest \
  /bin/kvmtop -c qemu:///system --printer text --cpu --mem
```

Notes for Docker: i) If libvirt is accessible via a local socket, a volume has to mount this socket inside the container. ii) Containers per design isolate the container in its own process namespace, which hinders reading from the proc filesystem. `--pid="host"` softens this isolation, so kvmtop has access to the hosts proc files. iii) the `--privileged` is required for the IO Collector, which reads files only accessible as root user.

## General Usage

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
  -b, --batch       [DEPRECATED: use --printer=text instead] use simple output
                    e.g. for scripts
  -p, --printer=    the output printer to use (valid printers: ncurses, text,
                    json) (default: ncurses)

Help Options:
  -h, --help        Show this help message
```

Exemplary output
```
UUID                                 name          cpu_cores cpu_total cpu_steal cpu_other_total cpu_other_steal
0dbe2ae8-1ee4-4b43-bdf3-b533dfe75486 ubuntu14.04-2 2         53        0         5               1
```

Please note: although the connection to libvirt may work remote (e.g. via ssh), kvmtop requires access to the /proc file system of the hypervisor's operating system. You can use the `--connection` to connect to a remote libvirt, but need to mount the remote proc fs and specify the location with `--procfs`.

### Printers and Outputs

Printers define the representation of the monitoring data. This can be for humans in ncurses, or for further processing text (space separated) or json.

Outputs define the location where the printers send data to. Output works for text and json printers, yet not for ncurses. The output may be a file or a remote tcp server.

Example scenarios:

```
# write monitoring data to log file
kvmtop --cpu --printer=text --output=file --target=/var/log/kvmtop.log

# send mointoring data to tcp server (e.g. logstash with tcp input)
kvmtop --cpu --printer=json --output=tcp --target=127.0.0.1:12345
```

## Collectors & Their Fields

| Collector | cli option | description |
| --- | --- | --- |
| CPU Collector | --cpu | CPU Stats (host and VMs) like cores, utilisation, frequency|
| Memory Collector | --mem | Memory stats (host and VMs)  like capacity, allocation, faults |
| Disk Collector | --disk | Disk stats (host and VMs) like capacity, utilisation, reads/writes, etc. |
| Network Collector | --net | Network stats (host and VMs) like transmitted and received bytes, packets, errors, etc. |
| I/O Collector | --io | Disk I/O stats (VMs) like reads/writes |

A more detailed list, including all metrics is available here at [./docs/README.md](https://github.com/cha87de/kvmtop/blob/master/docs/README.md).

## kvmtop with InfluxDB

kvmtop can be used as a monitoring agent to send data to an InfluxDB instance: kvmtop transmits JSON data via TCP to logstash, while logstash writes to InfluxDB. More detailes are available at [https://github.com/cha87de/kvmtop-datasink/](https://github.com/cha87de/kvmtop-datasink/).

```    
                                      kvmtop-datasink
                  +-----------------------------------------------------+
                  |                                                     |
+------------     | +------------+     +------------+     +-----------+ |
|           |     | |            |     |            |     |           | |
|  kvmtop   +---> | |  logstash  +---> |  influxdb  +---> |  grafana  | |
|           |     | |            |     |            |     |           | |
+------------     | +------------+     +------------+     +-----------+ |
                  |                                                     |
                  +-----------------------------------------------------+
```

# Development Guide

Install the golang binary and the required dependencies libvirt-dev and libncurses5-dev packages. Create a new folder as your kvmtop workspace, e.g. /opt/kvmtop.  Then follow these steps:
```
cd /opt/kvmtop
export GOPATH=$(pwd) # take workspace as GOPATH
go get -d github.com/cha87de/kvmtop/...   # download all source files, including depencencies
go install github.com/cha87de/kvmtop/...  # compile kvmtop binaries
```
The resulting binaries are then located in `/opt/kvmtop/bin`.

Further reading: https://golang.org/doc/code.html
