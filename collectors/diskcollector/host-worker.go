package diskcollector

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/util"

	"github.com/cha87de/kvmtop/models"
)

func diskHostLookup(host *models.Host) {

	/*
		// find relevant devices
		devices := []string{}
		mounts := util.GetProcMounts()
		diskSources := strings.Split(host.GetMetricString("disk_sources", 0), ",")
		for _, source := range diskSources {
			// find best matching mountpoint
			var bestMount util.ProcMount
			for _, mount := range mounts {
				// matches at all?
				if !strings.HasPrefix(source, mount.Mountpoint) {
					continue
				}
				// matches better than already found one?
				if len(bestMount.Mountpoint) < len(mount.Mountpoint) {
					bestMount = mount
				}
			}
			// add bestMount to devices, if not contained
			found := false
			for _, device := range devices {
				if device == bestMount.Device {
					found = true
					break
				}
			}
			if !found {
				device := filepath.Base(bestMount.Device)
				devices = append(devices, device)
			}
		}
	*/

	devices := []string{}
	if config.Options.StorageDevice != "" {
		devices = strings.Split(config.Options.StorageDevice, ",")
	}

	// lookup diskstats for relevant devices
	diskstats := util.GetProcDiskstats()
	combinedDiskstat := util.ProcDiskstat{}
	combinedDiskstatCounts := uint64(0)
	if len(devices) > 0 {
		// consider only relevant devices
		for _, device := range devices {
			if stats, ok := diskstats[device]; ok {
				combinedDiskstat.Reads += stats.Reads
				combinedDiskstat.ReadsMerged += stats.ReadsMerged
				combinedDiskstat.SectorsRead += stats.SectorsRead
				combinedDiskstat.TimeReading += stats.TimeReading
				combinedDiskstat.Writes += stats.Writes
				combinedDiskstat.WritesMerged += stats.WritesMerged
				combinedDiskstat.SectorsWritten += stats.SectorsWritten
				combinedDiskstat.TimeWriting += stats.TimeWriting
				combinedDiskstat.CurrentOps += stats.CurrentOps
				combinedDiskstat.TimeForOps += stats.TimeForOps
				combinedDiskstat.WeightedTimeForOps += stats.WeightedTimeForOps
				combinedDiskstatCounts++
			}
		}
	} else {

		// consider all available devices (clean duplicates like sda and sda1)
		diskstats = clearDuplicateDevices(diskstats)
		for _, stats := range diskstats {
			combinedDiskstat.Reads += stats.Reads
			combinedDiskstat.ReadsMerged += stats.ReadsMerged
			combinedDiskstat.SectorsRead += stats.SectorsRead
			combinedDiskstat.TimeReading += stats.TimeReading
			combinedDiskstat.Writes += stats.Writes
			combinedDiskstat.WritesMerged += stats.WritesMerged
			combinedDiskstat.SectorsWritten += stats.SectorsWritten
			combinedDiskstat.TimeWriting += stats.TimeWriting
			combinedDiskstat.CurrentOps += stats.CurrentOps
			combinedDiskstat.TimeForOps += stats.TimeForOps
			combinedDiskstat.WeightedTimeForOps += stats.WeightedTimeForOps
			combinedDiskstatCounts++
		}
	}

	host.AddMetricMeasurement("disk_device_reads", models.CreateMeasurement(combinedDiskstat.Reads))
	host.AddMetricMeasurement("disk_device_readsmerged", models.CreateMeasurement(combinedDiskstat.ReadsMerged))
	host.AddMetricMeasurement("disk_device_sectorsread", models.CreateMeasurement(combinedDiskstat.SectorsRead))
	host.AddMetricMeasurement("disk_device_timereading", models.CreateMeasurement(combinedDiskstat.TimeReading))
	host.AddMetricMeasurement("disk_device_writes", models.CreateMeasurement(combinedDiskstat.Writes))
	host.AddMetricMeasurement("disk_device_writesmerged", models.CreateMeasurement(combinedDiskstat.WritesMerged))
	host.AddMetricMeasurement("disk_device_sectorswritten", models.CreateMeasurement(combinedDiskstat.SectorsWritten))
	host.AddMetricMeasurement("disk_device_timewriting", models.CreateMeasurement(combinedDiskstat.TimeWriting))
	host.AddMetricMeasurement("disk_device_currentops", models.CreateMeasurement(combinedDiskstat.CurrentOps))
	host.AddMetricMeasurement("disk_device_timeforops", models.CreateMeasurement(combinedDiskstat.TimeForOps))
	host.AddMetricMeasurement("disk_device_weightedtimeforops", models.CreateMeasurement(combinedDiskstat.WeightedTimeForOps))
	host.AddMetricMeasurement("disk_device_count", models.CreateMeasurement(combinedDiskstatCounts))
}

func diskHostCollect(host *models.Host) {

}

func diskPrintHost(host *models.Host) []string {
	diskDeviceReads := host.GetMetricDiffUint64("disk_device_reads", true)
	diskDeviceReadsmerged := host.GetMetricDiffUint64("disk_device_readsmerged", true)
	diskDeviceSectorsread := host.GetMetricDiffUint64("disk_device_sectorsread", true)
	diskDeviceTimereading := host.GetMetricDiffUint64("disk_device_timereading", true)
	diskDeviceWrites := host.GetMetricDiffUint64("disk_device_writes", true)
	diskDeviceWritesmerged := host.GetMetricDiffUint64("disk_device_writesmerged", true)
	diskDeviceSectorswritten := host.GetMetricDiffUint64("disk_device_sectorswritten", true)
	diskDeviceTimewriting := host.GetMetricDiffUint64("disk_device_timewriting", true)
	diskDeviceCurrentops := host.GetMetricDiffUint64("disk_device_currentops", true)
	diskDeviceTimeforops := host.GetMetricDiffUint64("disk_device_timeforops", true)
	diskDeviceWeightedtimeforops := host.GetMetricDiffUint64("disk_device_weightedtimeforops", true)
	//diskDeviceCountStr, _ := host.GetMetricUint64("disk_device_count", 1)
	//diskDeviceCount, _ := strconv.Atoi(diskDeviceCountStr)

	ioutil := timeToPercent(host, "disk_device_timeforops")

	result := append([]string{diskDeviceReads}, diskDeviceWrites, ioutil)
	if config.Options.Verbose {
		result = append(result, diskDeviceReadsmerged, diskDeviceSectorsread, diskDeviceTimereading)
		result = append(result, diskDeviceWritesmerged, diskDeviceSectorswritten, diskDeviceTimewriting, diskDeviceCurrentops, diskDeviceTimeforops)
		result = append(result, diskDeviceWeightedtimeforops)
	}

	return result
}

func timeToPercent(host *models.Host, metricName string) string {
	var output string
	var percent float64
	if metric, ok := host.GetMetric(metricName); ok {
		if len(metric.Values) >= 2 {
			// get first value
			byteValue1 := metric.Values[0].Value
			reader1 := bytes.NewReader(byteValue1)
			decoder1 := gob.NewDecoder(reader1)
			var value1 uint64
			decoder1.Decode(&value1)

			// get second value
			byteValue2 := metric.Values[1].Value
			reader2 := bytes.NewReader(byteValue2)
			decoder2 := gob.NewDecoder(reader2)
			var value2 uint64
			decoder2.Decode(&value2)

			// calculate value diff per time
			value := float64(value1 - value2)

			// get time diff
			ts1 := metric.Values[0].Timestamp
			ts2 := metric.Values[1].Timestamp
			diffSeconds := ts1.Sub(ts2).Seconds()
			valuePerSecond := value / 1000 // since value is in ms
			ratio := valuePerSecond / diffSeconds
			percent = ratio * 100 // compute it as percent

			output = fmt.Sprintf("%.0f", percent)
		}
	}
	return output

}

func clearDuplicateDevices(diskstats map[string]util.ProcDiskstat) map[string]util.ProcDiskstat {
	result := make(map[string]util.ProcDiskstat)
	keys := make([]string, 0, len(diskstats))
	for k := range diskstats {
		keys = append(keys, k)
	}

	// remove duplicates like sda and sda1 - only consider sda1
	for key, stats := range diskstats {
		// is there a key in keys which is longer?
		considerDisk := true
		for _, k := range keys {
			if strings.HasPrefix(k, key) && len(k) > len(key) {
				// found more detailed device name
				considerDisk = false
				break
			}
		}
		if considerDisk {
			result[key] = stats
		}
	}
	return result

}
