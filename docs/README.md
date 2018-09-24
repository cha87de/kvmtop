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
| - | - | - |

missing: total bandwidth, overall utilisation, errors, losses, queue sizes, ?

### Network Virtual Machine Metrics

| Metric | Source | Description |
| --- | --- | --- |
| net_ifs | libvirt | |
| net_tx | libvirt | |
| net_rx | libvirt | |

missing: queue sizes, errors, losses, ?

TODO: change source to proc!

## Disk Collector

Enable with parameter `--disk`.

### Disk Host Metrics

| Metric | Source | Description |
| --- | --- | --- |
| disk_sources ? | | |
| disk_device_reads | proc | |
| disk_device_readsmerged | proc | |
| disk_device_sectorsread | proc | |
| disk_device_timereading | proc | |
| disk_device_writes | proc | |
| disk_device_writesmerged | proc | |
| disk_device_sectorswritten | proc | |
| disk_device_timewriting | proc | |
| disk_device_currentops | proc | |
| disk_device_timeforops | proc | |
| disk_device_weightedtimeforops | proc | |

missing: capacity, used capacity, fs cache misses, disk scheduler infos?, max bandwidth, ?

### Disk Virtual Machine Metrics

| Metric | Source | Description |
| --- | --- | --- |
| disks_disks | libvirt | |
| disk_stats_errs | libvirt | |
| disk_stats_flushreq | libvirt | |
| disk_stats_flushtotaltimes | libvirt | |
| disk_stats_rdbytes | libvirt | |
| disk_stats_rdreq | libvirt | |
| disk_stats_rdtotaltimes | libvirt | |
| disk_stats_wrbytes | libvirt | |
| disk_stats_wrreq | libvirt | |
| disk_stats_wrtotaltimes | libvirt | |
| disk_delayblkio | proc | |

missing: capacity, used capacity, ?

## IO Collector

TODO:  explain the difference to Disk collector.

Enable with parameter `--io`.

*Please note:* this collector requires root access to /proc.

### IO Host Metrics

| Metric | Source | Description |
| --- | --- | --- |
| - | - | - |

### IO Virtual Machine Metrics

| Metric | Source | Description |
| --- | --- | --- |
| io_rchar | proc | |
| io_wchar | proc | |
| io_syscr | proc | |
| io_syscw | proc | |
| io_read_bytes | proc | |
| io_write_bytes | proc | |
| io_cancelled_write_bytes | proc | |