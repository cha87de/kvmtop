# Metric Collectors

Available metrics, their source and description.

## CPU Collector

Enable with parameter `--cpu`.

### CPU Host Metrics

| Metric | Source | Description |
| --- | --- | --- |
| cpu_meanfreq | proc | |
| cpu_cores | proc | |

missing: architecture

### CPU Virtual Machine Metrics

| Metric | Source | Description |
| --- | --- | --- |
| cpu_cores | proc | |
| cpu_total | proc | |
| cpu_steal | proc | |
| cpu_other_total | proc | |
| cpu_other_steal | proc | |

missing: none

## Memory Collector

Enable with parameter `--mem`.

### Memory Host Metrics

| Metric | Source | Description |
| --- | --- | --- |
| - | - | - |

missing: total, used, faults, frequency, ?

### Memory Virtual Machine Metrics

| Metric | Source | Description |
| --- | --- | --- |
| ram_total | libvirt | |
| ram_used | libvirt | |
| ram_vsize | proc | |
| ram_rss | proc | |
| ram_minflt | proc | |
| ram_cminflt | proc | |
| ram_majflt | proc | |
| ram_cmajflt | proc | |

missing: total, used, faults, frequency, ?

## Network Collector

Enable with parameter `--net`.

### Network Host Metrics

| Metric | Source | Description |
| --- | --- | --- |
| net_host_receivedBytes| proc | |
| net_host_receivedPackets| proc | |
| net_host_receivedErrs| proc | |
| net_host_receivedDrop| proc | |
| net_host_receivedFifo| proc | |
| net_host_receivedFrame| proc | |
| net_host_receivedCompressed| proc | |
| net_host_receivedMulticast| proc | |
| net_host_transmittedBytes| proc | |
| net_host_transmittedPackets| proc | |
| net_host_transmittedErrs| proc | |
| net_host_transmittedDrop| proc | |
| net_host_transmittedFifo| proc | |
| net_host_transmittedColls| proc | |
| net_host_transmittedCarrier| proc | |
| net_host_transmittedCompressed| proc | |

missing: total bandwidth, queue sizes, ?

### Network Virtual Machine Metrics

| Metric | Source | Description |
| --- | --- | --- |
| net_receivedBytes| proc | |
| net_receivedPackets| proc | |
| net_receivedErrs| proc | |
| net_receivedDrop| proc | |
| net_receivedFifo| proc | |
| net_receivedFrame| proc | |
| net_receivedCompressed| proc | |
| net_receivedMulticast| proc | |
| net_transmittedBytes| proc | |
| net_transmittedPackets| proc | |
| net_transmittedErrs| proc | |
| net_transmittedDrop| proc | |
| net_transmittedFifo| proc | |
| net_transmittedColls| proc | |
| net_transmittedCarrier| proc | |
| net_transmittedCompressed| proc | |

missing: wait times / latencies, queue sizes, ?

## Disk Collector

Enable with parameter `--disk`.

### Disk Host Metrics

| Metric | Source | Description |
| --- | --- | --- |
| disk_sources ? | | |
| disk_device_reads | proc | reads completed successfully |
| disk_device_readsmerged | proc | reads merged |
| disk_device_sectorsread | proc | sectors read |
| disk_device_timereading | proc | time spent reading (ms) |
| disk_device_writes | proc | writes completed |
| disk_device_writesmerged | proc | writes merged |
| disk_device_sectorswritten | proc | sectors written |
| disk_device_timewriting | proc | time spent writing (ms) |
| disk_device_currentops | proc | I/Os currently in progress |
| disk_device_timeforops | proc | time spent doing I/Os (ms) |
| disk_device_weightedtimeforops | proc | weighted time spent doing I/Os (ms) |

missing: capacity, used capacity, fs cache misses, disk scheduler infos?, max bandwidth, ?

### Disk Virtual Machine Metrics

| Metric | Source | Description |
| --- | --- | --- |
| disks_disks | libvirt | |
| disk_stats_errs | libvirt | only for Xen Hypervisor - needs to be removed |
| disk_stats_flushreq | libvirt | represents the total flush requests of the block device |
| disk_stats_flushtotaltimes | libvirt | represents the total time spend on cache flushing in nano-seconds of the block device |
| disk_stats_rdbytes | libvirt | represents the total number of read bytes of the block device |
| disk_stats_rdreq | libvirt | represents the total read requests of the block device |
| disk_stats_rdtotaltimes | libvirt | represents the total time spend on cache reads in nano-seconds of the block device |
| disk_stats_wrbytes | libvirt | represents the total number of write bytes of the block device |
| disk_stats_wrreq | libvirt | represents the total write requests of the block device |
| disk_stats_wrtotaltimes | libvirt | represents the total time spend on cache writes in nano-seconds of the block device |
| disk_delayblkio | proc | aggregated block I/O delays, measured in clock ticks (centiseconds) |

missing: capacity, used capacity, ?

**TODO:** change source to proc! see issue #8

## IO Collector

**TODO:** explain the difference to Disk collector.

Enable with parameter `--io`.

*Please note:* this collector requires root access to /proc.

### IO Host Metrics

| Metric | Source | Description |
| --- | --- | --- |
| - | - | - |

### IO Virtual Machine Metrics

| Metric | Source | Description |
| --- | --- | --- |
| io_rchar | proc | number of bytes the process read, using any read-like system call (from files, pipes, tty...) |
| io_wchar | proc | number of bytes the process wrote using any write-like system call |
| io_syscr | proc | number of read-like system call invocations that the process performed |
| io_syscw | proc | number of write-like system call invocations that the process performed |
| io_read_bytes | proc | number of bytes the process directly read from disk |
| io_write_bytes | proc | number of bytes the process originally dirtied in the page-cache (assuming they will go to disk later) |
| io_cancelled_write_bytes | proc | number of bytes the process "un-dirtied" - e.g. using an "ftruncate" call that truncated pages from the page-cache |
