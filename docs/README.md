# Metric Collectors

This documentation describes available metrics, their sources and description.

Metrics below the **verbose mode** sign are only available when using the
`--verbose` option. Metrics below the **internal metrics** sign are not exported
but only available internally to process the exported metrics.

Before any collector can run, kvmtop looks up active virtual machines from
libvirt and looks up processes from the proc filesystem by reading the
directories from the proc filesystem. From libvirt, the UUID and the name of
virtual machines are collected. With the UUID, the correlated (parent) process
is searched, by using the proc/${pid}/cmdline file and matching for the UUID.
These steps look up the basic details of the virtual machines (UUID, name,
parent process id), which are necessary for any of the collectors to work.

## CPU Collector

Enable with parameter `--cpu`.

### CPU Host Metrics

| Metric           | Source                       | Description                                                     | Cycle   |
|------------------|------------------------------|-----------------------------------------------------------------|---------|
| cpu_cores        | proc/cpuinfo                 | Total number of CPU cores                                       | lookup  |
| cpu_curfreq      | /sys/devices/system/cpu/cpu* | Current mean frequency over all CPU cores                       | lookup  |
| cpu_user         | proc/stat                    | utilisation in percent spent in user space                      | collect |
| cpu_system       | proc/stat                    | utilisation in percent spent in system space                    | collect |
| cpu_idle         | proc/stat                    | utilisation in percent spent in idle                            | collect |
| cpu_steal        | proc/stat                    | utilisation in percent stolen due to overbooking                | collect |
| **verbose mode** |                              |                                                                 |         |
| cpu_minfreq      | /sys/devices/system/cpu/cpu* | Minimal frequency the CPU can run on                            | lookup  |
| cpu_maxfreq      | /sys/devices/system/cpu/cpu* | Maximum frequency the CPU can run on                            | lookup  |
| cpu_nice         | proc/stat                    | utilisation in percent spent in nice mode                       | collect |
| cpu_iowait       | proc/stat                    | utilisation in percent spent waiting for IO                     | collect |
| cpu_irq          | proc/stat                    | time-relative amount of interrupts                              | collect |
| cpu_softirq      | proc/stat                    | time-relative amount of soft interrupts                         | collect |
| cpu_guest        | proc/stat                    | utilisation in percent spent in guest (VM) space                | collect |
| cpu_guestnice    | proc/stat                    | utilisation in percent spent in guest (VM) space with nice mode | collect |

missing: architecture

### CPU Virtual Machine Metrics

The lookup cycle first reads process IDs from libvirt, then collects process IDs
from the proc filesystem. From both information, the cpu usage for virtual CPU
cores and for overhead can be differentiated (processes which are not listed by
libvirt but assigned to the parent process ID of the virtual machine). The proc
filesystem then provides utilisation values and steal values from cpu cycle
queueing from the cpu scheduler statistics, per process.

| Metric                     | Source                | Description                                                                | Cycle   |
|----------------------------|-----------------------|----------------------------------------------------------------------------|---------|
| cpu_cores                  | libvirt               | Total number of virtual CPU cores                                          | lookup  |
| cpu_total                  | internal metrics      | utilisation in percent over all virtual CPU cores as experienced inside VM | collect |
| cpu_steal                  | internal metrics      | stolen (reduced) cpu utilisation over all virtual CPU cores                | collect |
| **verbose mode**           |                       |                                                                            |         |
| cpu_other_total            | internal metrics      | additional utilisation in percent (runtime overhead) of VM                 | collect |
| cpu_other_steal            | internal metrics      | additional stolen utilisation of VM                                        | collect |
| **internal metrics**       |                       |                                                                            |         |
| cpu_threadIDs              | libvirt + proc        | list of process IDs assigned to virtual CPU cores of a VM                  | lookup  |
| cpu_otherThreadIDs         | libvirt + proc        | list of process IDs assigned to VM                                         | lookup  |
| cpu_times_${pid}           | proc/${pid}/schedstat | time (counter) spent processing virtual CPU cores                          | collect |
| cpu_runqueues_${pid}       | proc/${pid}/schedstat | time (counter) spent waiting for processing for virtual CPU cores          | collect |
| cpu_other_times_${pid}     | proc/${pid}/schedstat | time (conter) spent for processing runtime overhead                        | collect |
| cpu_other_runqueues_${pid} | proc/${pid}/schedstat | time (counter) spent for waiting on processing for runtime overhead        | collect |


missing: none

## Memory Collector

Enable with parameter `--mem`.

### Memory Host Metrics

| Metric                | Source       | Description                                         | Cycle   |
|-----------------------|--------------|-----------------------------------------------------|---------|
| ram_Total             | proc/meminfo | Total amount of physical ram available on host      | collect |
| ram_Free              | proc/meminfo | Total amount of free physical ram available on host | collect |
| ram_Available         | proc/meminfo | Total amount of available (used) ram on the host    | collect |
| **verbose mode**      |              |                                                     |         |
| ram_Buffers           | proc/meminfo |                                                     | collect |
| ram_Cached            | proc/meminfo |                                                     | collect |
| ram_SwapCached        | proc/meminfo |                                                     | collect |
| ram_Active            | proc/meminfo |                                                     | collect |
| ram_Inactive          | proc/meminfo |                                                     | collect |
| ram_ActiveAanon       | proc/meminfo |                                                     | collect |
| ram_InactiveAanon     | proc/meminfo |                                                     | collect |
| ram_ActiveFile        | proc/meminfo |                                                     | collect |
| ram_InactiveFile      | proc/meminfo |                                                     | collect |
| ram_Unevictable       | proc/meminfo |                                                     | collect |
| ram_Mlocked           | proc/meminfo |                                                     | collect |
| ram_SwapTotal         | proc/meminfo |                                                     | collect |
| ram_SwapFree          | proc/meminfo |                                                     | collect |
| ram_Dirty             | proc/meminfo |                                                     | collect |
| ram_Writeback         | proc/meminfo |                                                     | collect |
| ram_AnonPages         | proc/meminfo |                                                     | collect |
| ram_Mapped            | proc/meminfo |                                                     | collect |
| ram_Shmem             | proc/meminfo |                                                     | collect |
| ram_Slab              | proc/meminfo |                                                     | collect |
| ram_SReclaimable      | proc/meminfo |                                                     | collect |
| ram_SUnreclaim        | proc/meminfo |                                                     | collect |
| ram_KernelStack       | proc/meminfo |                                                     | collect |
| ram_PageTables        | proc/meminfo |                                                     | collect |
| ram_NFSUnstable       | proc/meminfo |                                                     | collect |
| ram_Bounce            | proc/meminfo |                                                     | collect |
| ram_WritebackTmp      | proc/meminfo |                                                     | collect |
| ram_CommitLimit       | proc/meminfo |                                                     | collect |
| ram_CommittedAS       | proc/meminfo |                                                     | collect |
| ram_VmallocTotal      | proc/meminfo |                                                     | collect |
| ram_VmallocUsed       | proc/meminfo |                                                     | collect |
| ram_VmallocChunk      | proc/meminfo |                                                     | collect |
| ram_HardwareCorrupted | proc/meminfo |                                                     | collect |
| ram_AnonHugePages     | proc/meminfo |                                                     | collect |
| ram_ShmemHugePages    | proc/meminfo |                                                     | collect |
| ram_ShmemPmdMapped    | proc/meminfo |                                                     | collect |
| ram_HugePagesTotal    | proc/meminfo |                                                     | collect |
| ram_HugePagesFree     | proc/meminfo |                                                     | collect |
| ram_HugePagesRsvd     | proc/meminfo |                                                     | collect |
| ram_HugePagesSurp     | proc/meminfo |                                                     | collect |
| ram_Hugepagesize      | proc/meminfo |                                                     | collect |
| ram_Hugetlb           | proc/meminfo |                                                     | collect |
| ram_DirectMap4k       | proc/meminfo |                                                     | collect |
| ram_DirectMap2M       | proc/meminfo |                                                     | collect |
| ram_DirectMap1G       | proc/meminfo |                                                     | collect |

missing: frequency

### Memory Virtual Machine Metrics

| Metric           | Source           | Description                                        | Cycle   |
|------------------|------------------|----------------------------------------------------|---------|
| ram_total        | libvirt          | Maximum amount of memory the VM can allocate       | lookup  |
| ram_used         | libvirt          | Allocated amount of memory the VM claimed          | lookup  |
| **verbose mode** |                  |                                                    |         |
| ram_vsize        | proc/${pid}/stat | Virtual memory the VM allocates on the host        | collect |
| ram_rss          | proc/${pid}/stat | Actual memory the VM allocated in the hosts memory | collect |
| ram_minflt       | proc/${pid}/stat | Minor faults where the VM was affected             | collect |
| ram_cminflt      | proc/${pid}/stat | Minor faults with child processes                  | collect |
| ram_majflt       | proc/${pid}/stat | Major faults where the VM was affected             | collect |
| ram_cmajflt      | proc/${pid}/stat | Major faults with child processes                  | collect |

missing: none

## Network Collector

Enable with parameter `--net`.

First, libvirt is used to identify relevant virtual and physical network
interfaces. With these unique network interface names the proc and sys
filesystems are queried to retrieve relevant metrics for the host and per
virtual machine. The separation, even if network interfaces are used by multiple
virtual machines, is possible, since the proc filesystem allows to query the
overall statistics but also with a limited view for a single process.

### Network Host Metrics

| Metric                         | Source                          | Description                                                    | Cycle   |
|--------------------------------|---------------------------------|----------------------------------------------------------------|---------|
| net_host_receivedBytes         | proc/net/dev                    | Number of received bytes (sum over all relevant interfaces)    | collect |
| net_host_transmittedBytes      | proc/net/dev                    | Number of transmitted bytes (sum over all relevant interfaces) | collect |
| net_host_speed                 | /sys/class/net/${devName}/speed | Network devices' maximum speed                                 | lookup  |
| **verbose mode**               |                                 |                                                                |         |
| net_host_receivedPackets       | proc/net/dev                    | Number of received packets                                     | collect |
| net_host_receivedErrs          | proc/net/dev                    | Number of errors of received packets                           | collect |
| net_host_receivedDrop          | proc/net/dev                    | Number of drops of received packets                            | collect |
| net_host_receivedFifo          | proc/net/dev                    | Number of FIFO buffer errors for received packets              | collect |
| net_host_receivedFrame         | proc/net/dev                    | Number of framing errors for received packets                  | collect |
| net_host_receivedCompressed    | proc/net/dev                    | Number of compressed packets received by device driver         | collect |
| net_host_receivedMulticast     | proc/net/dev                    | Number of received multicast frames                            | collect |
| net_host_transmittedPackets    | proc/net/dev                    | Number of transmitted packets                                  | collect |
| net_host_transmittedErrs       | proc/net/dev                    | Number of errors of transmitted packets                        | collect |
| net_host_transmittedDrop       | proc/net/dev                    | Number of drops of transmitted packets                         | collect |
| net_host_transmittedFifo       | proc/net/dev                    | Number of FIFO buffer errors for transmitted packets           | collect |
| net_host_transmittedColls      | proc/net/dev                    | Number of detected packet collisions while transmitting        | collect |
| net_host_transmittedCarrier    | proc/net/dev                    | Number of detected carrier losses while transmitting           | collect |
| net_host_transmittedCompressed | proc/net/dev                    | Number of compressed packets transmitted by device driver      | collect |
| **internal metrics**           |                                 |                                                                |         |
| net_host_ifs                   | libvirt                         | List of relevant physical network interfaces                   | lookup  |

missing: queue sizes

### Network Virtual Machine Metrics

| Metric                    | Source              | Description                                               | Cycle   |
|---------------------------|---------------------|-----------------------------------------------------------|---------|
| net_receivedBytes         | proc/${pid}/net/dev | Number of received bytes                                  | collect |
| net_transmittedBytes      | proc/${pid}/net/dev | Number of transmitted bytes                               | collect |
| **verbose mode**          |                     |                                                           |         |
| net_receivedPackets       | proc/${pid}/net/dev | Number of received packets                                | collect |
| net_receivedErrs          | proc/${pid}/net/dev | Number of errors of received packets                      | collect |
| net_receivedDrop          | proc/${pid}/net/dev | Number of drops of received packets                       | collect |
| net_receivedFifo          | proc/${pid}/net/dev | Number of FIFO buffer errors for received packets         | collect |
| net_receivedFrame         | proc/${pid}/net/dev | Number of framing errors for received packets             | collect |
| net_receivedCompressed    | proc/${pid}/net/dev | Number of compressed packets received by device driver    | collect |
| net_receivedMulticast     | proc/${pid}/net/dev | Number of received multicast frames                       | collect |
| net_transmittedPackets    | proc/${pid}/net/dev | Number of transmitted packets                             | collect |
| net_transmittedErrs       | proc/${pid}/net/dev | Number of errors of transmitted packets                   | collect |
| net_transmittedDrop       | proc/${pid}/net/dev | Number of drops of transmitted packets                    | collect |
| net_transmittedFifo       | proc/${pid}/net/dev | Number of FIFO buffer errors for transmitted packets      | collect |
| net_transmittedColls      | proc/${pid}/net/dev | Number of detected packet collisions while transmitting   | collect |
| net_transmittedCarrier    | proc/${pid}/net/dev | Number of detected carrier losses while transmitting      | collect |
| net_transmittedCompressed | proc/${pid}/net/dev | Number of compressed packets transmitted by device driver | collect |
| **internal metrics**      |                     |                                                           |         |
| net_interfaces            | libvirt             | List of (virtual) network interfaces                      | lookup  |

missing: queue sizes

## Disk Collector

Enable with parameter `--disk`.

The disk collector first looks up block storage devices for virtual machines, to
then combine a list of disk_sources for each VM and a total list for the host.
kvmtop further allows via the configuration parameter `--storedev` to manually
specify the hosts storage devices, which are then used instead to query statistics.

The metrics ioutil, queue size, queue time and service time are combined
metrics, which are dynamically calculated from other directly collectable
metrics. The ioutil is calculated by using the total time during which I/Os were
in progress, divided by the sampling interval. The queue size is the weighted
number of milliseconds spent doing I/Os divided by the milliseconds elapsed
between two measurements. The queue time is the weighted time spent doing I/Os,
divided by the sum of reads, readsMerged, writes, writesMerged, and currentOps.
The service time is the weighted time spent doing I/Os, divided by the sum of
reads, readsMerged, writes, and writesMerged.


### Disk Host Metrics

| Metric                         | Source         | Description                                    | Cycle   |
|--------------------------------|----------------|------------------------------------------------|---------|
| disk_device_reads              | proc/diskstats | Reads completed successfully                   | lookup  |
| disk_device_writes             | proc/diskstats | Writes completed successfully                  | lookup  |
| disk_device_ioutil             | combined       | Saturation of handling I/O requests in percent | collect |
| **verbose mode**               |                |                                                |         |
| disk_device_readsmerged        | proc/diskstats | Reads merged                                   | lookup  |
| disk_device_sectorsread        | proc/diskstats | Sectors read                                   | lookup  |
| disk_device_timereading        | proc/diskstats | Time spent reading (ms)                        | lookup  |
| disk_device_writesmerged       | proc/diskstats | Writes merged                                  | lookup  |
| disk_device_sectorswritten     | proc/diskstats | Sectors written                                | lookup  |
| disk_device_timewriting        | proc/diskstats | Time spent writing (ms)                        | lookup  |
| disk_device_currentops         | proc/diskstats | I/Os currently in progress                     | lookup  |
| disk_device_timeforops         | proc/diskstats | Time spent doing I/Os (ms)                     | lookup  |
| disk_device_weightedtimeforops | proc/diskstats | Weighted time spent doing I/Os (ms)            | lookup  |
| disk_device_count              | proc/diskstats | Number of relevant disks                       | lookup  |
| disk_device_queuesize          | combined       | Numer of queued I/O requests                   | collect |
| disk_device_queuetime          | combined       | The average queue time or I/O requests         | collect |
| disk_device_servicetime        | combined       | The average time serving I/O requests          | collect |
| **internal metrics**           |                |                                                |         |
| disk_sources                   | libvirt        | List of relevant disks                         | lookup  |

missing: capacity, used capacity, fs cache misses, disk scheduler infos?, max
bandwidth

### Disk Virtual Machine Metrics

Since the Linux kernel has no option to query disk statistics per process, the
virtual storage driver has to be queried to collect the metrics like requests or
bytes read / written to the virtual storage. This storage driver is part of the
QEMU emulator, which can be queried via libvirt.

The virtual machine ioutil is an estimated value, based on the host's ioutil
value and the relative read and write requests of this virtual machine to others
and the overall host requests.

| Metric                     | Source           | Description                                                                        | Cycle   |
|----------------------------|------------------|------------------------------------------------------------------------------------|---------|
| disk_size_capacity         | libvirt          | Maximum capacity of the virtual block devices (sum if multiple devs.)              | lookup  |
| disk_size_allocation       | libvirt          | Allocated space of the virtual block devices (sum if multiple devs.)               | lookup  |
| disk_ioutil                | combined         | Estimated I/O utilisation for the VM in percent                                    | collect |
| **verbose mode**           |                  |                                                                                    |         |
| disk_size_physical         | libvirt          | Physical space required to serve the virtual block devices (sum if multiple devs.) | lookup  |
| disk_stats_flushreq        | libvirt          | Flush requests of the block device                                                 | lookup  |
| disk_stats_flushtotaltimes | libvirt          | Time spend on cache flushing in nano-seconds of the block device                   | lookup  |
| disk_stats_rdbytes         | libvirt          | Number of read bytes of the block device                                           | lookup  |
| disk_stats_rdreq           | libvirt          | Read requests of the block device                                                  | lookup  |
| disk_stats_rdtotaltimes    | libvirt          | Time spend on cache reads in nano-seconds of the block device                      | lookup  |
| disk_stats_wrbytes         | libvirt          | Number of write bytes of the block device                                          | lookup  |
| disk_stats_wrreq           | libvirt          | Write requests of the block device                                                 | lookup  |
| disk_stats_wrtotaltimes    | libvirt          | Time spend on cache writes in nano-seconds of the block device                     | lookup  |
| disk_delayblkio            | proc/${pid}/stat | Aggregated block I/O delays, measured in clock ticks (centiseconds)                | collect |
| **internal metrics**       |                  |                                                                                    |         |
| disk_sources               | libvirt          | List of (virtual) disk sources                                                     | lookup  |

## IO Collector

Enable with parameter `--io`.

The IO Collector extends the disk collector with utilisation metrics from the
proc fs instead from libvirt. *Please note: this collector requires root access
to /proc on most Linux distributions.*

### IO Host Metrics

| Metric               | Source | Description | Cycle |
|----------------------|--------|-------------|-------|
| no metrics available |        |             |       |

### IO Virtual Machine Metrics

| Metric                   | Source         | Description                                                                                              | Cycle   |
|--------------------------|----------------|----------------------------------------------------------------------------------------------------------|---------|
| io_read_bytes            | proc/${pid}/io | Bytes the process directly read from disk                                                                | collect |
| io_write_bytes           | proc/${pid}/io | Bytes the process originally dirtied in the page-cache (assuming they will go to disk later)             | collect |
| **verbose mode**         |                |                                                                                                          | collect |
| io_rchar                 | proc/${pid}/io | Bytes the process read, using any read-like system call (from files, pipes, tty...)                      | collect |
| io_wchar                 | proc/${pid}/io | Bytes the process wrote using any write-like system call                                                 | collect |
| io_syscr                 | proc/${pid}/io | Read-like system call invocations that the process performed                                             | collect |
| io_syscw                 | proc/${pid}/io | Write-like system call invocations that the process performed                                            | collect |
| io_cancelled_write_bytes | proc/${pid}/io | Bytes the process "un-dirtied" - e.g. using an "ftruncate" call that truncated pages from the page-cache | collect |


## Host Collector

Enable with parameter `--host´.

The Host Collector extends the metrics by host specific details to identify the
current host where kvmtop is running on.

### Host Host Metrics

| Metric           | Source                                   | Description                                    | Cycle  |
|------------------|------------------------------------------|------------------------------------------------|--------|
| host_name        | /proc/sys/kernel/hostname                | The host name of the current host              | lookup |
| **verbose mode** |                                          |                                                |        |
| host_uuid        | /sys/devices/virtual/dmi/id/product_uuid | The hosts DMI UUID (requires root privileges!) | lookup |

### Host Virtual Machine Metrics

| Metric    | Source                    | Description                                          | Cycle  |
|-----------|---------------------------|------------------------------------------------------|--------|
| host_name | /proc/sys/kernel/hostname | The host name of the hypervisor the VM is running on | lookup |


## PSI Collector

Enable with parameter `--psi.

The PSI Collector extends the metrics by providing the Pressure Stall
Information (PSI) data from the host. These values indicate a resource
shortcoming in advance, before it actually occurs, for CPU, IO and memory.
Pressure Stall Information (PSI) controller works only with kernel 4.20 and
higher.

cf. https://facebookmicrosites.github.io/psi/docs/overview#pressure-metric-definitions

### Host Host Metrics

| Metric              | Source             | Description                                                                                     | Cycle   |
|---------------------|--------------------|-------------------------------------------------------------------------------------------------|---------|
| psi_some_cpu_avg60  | /proc/pressure/cpu | Ratio (percent) of time some (>= 1) tasks were delayed due to lack of CPU within 60 seconds     | collect |
| psi_some_io_avg60   | /proc/pressure/io  | Ratio (percent) of time some (>= 1) tasks were delayed due to lack of I/O within 60 seconds     | collect |
| psi_full_io_avg60   | /proc/pressure/io  | Ratio (percent) of time all tasks were delayed due to lack of I/O within 60 seconds             | collect |
| psi_some_mem_avg60  | /proc/pressure/mem | Ratio (percent) of time some (>= 1) tasks were delayed due to lack of memory within 60 seconds  | collect |
| psi_full_mem_avg60  | /proc/pressure/mem | Ratio (percent) of time all tasks were delayed due to lack of memory within 60 seconds          | collect |
| **verbose mode**    |                    |                                                                                                 |         |
| psi_some_cpu_avg10  | /proc/pressure/cpu | Ratio (percent) of time some (>= 1) tasks were delayed due to lack of CPU within 10 seconds     | collect |
| psi_some_cpu_avg300 | /proc/pressure/cpu | Ratio (percent) of time some (>= 1) tasks were delayed due to lack of CPU within 300 seconds    | collect |
| psi_some_cpu_total  | /proc/pressure/cpu | Total absolute delay time for some (>= 1) tasks (in microseconds) for CPU                       | collect |
| psi_some_io_avg10   | /proc/pressure/io  | Ratio (percent) of time some (>= 1) tasks were delayed due to lack of I/O within 10 seconds     | collect |
| psi_some_io_avg300  | /proc/pressure/io  | Ratio (percent) of time some (>= 1) tasks were delayed due to lack of I/O within 300 seconds    | collect |
| psi_some_io_total   | /proc/pressure/io  | Total absolute delay time for some (>= 1) tasks (in microseconds) for I/O                       | collect |
| psi_full_io_avg10   | /proc/pressure/io  | Ratio (percent) of time all tasks were delayed due to lack of I/O within 10 seconds             | collect |
| psi_full_io_avg300  | /proc/pressure/io  | Ratio (percent) of time all tasks were delayed due to lack of I/O within 300 seconds            | collect |
| psi_full_io_total   | /proc/pressure/io  | Total absolute delay time for all tasks (in microseconds) for I/O                               | collect |
| psi_some_mem_avg10  | /proc/pressure/mem | Ratio (percent) of time some (>= 1) tasks were delayed due to lack of memory within 10 seconds  | collect |
| psi_some_mem_avg300 | /proc/pressure/mem | Ratio (percent) of time some (>= 1) tasks were delayed due to lack of memory within 300 seconds | collect |
| psi_some_mem_total  | /proc/pressure/mem | Total absolute delay time for some (>= 1) tasks (in microseconds) for memory                    | collect |
| psi_full_mem_avg10  | /proc/pressure/mem | Ratio (percent) of time all tasks were delayed due to lack of memory within 10 seconds          | collect |
| psi_full_mem_avg300 | /proc/pressure/mem | Ratio (percent) of time all tasks were delayed due to lack of memory within 300 seconds         | collect |
| psi_full_mem_total  | /proc/pressure/mem | Total absolute delay time for all tasks (in microseconds) for memory                            | collect |

### Host Virtual Machine Metrics

| Metric               | Source | Description | Cycle |
|----------------------|--------|-------------|-------|
| no metrics available |        |             |       |