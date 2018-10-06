package diskcollector

import (
	"github.com/cha87de/kvmtop/collectors"
	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/util"

	"github.com/cha87de/kvmtop/models"
)

func diskHostLookup(host *models.Host) {

	/*
		// find relevant devices
		devices := []string{}
		mounts := util.GetProcMounts()
		diskSources := strings.Split(collectors.GetMetricString(host.Measurable, "disk_sources", 0), ",")
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
	// lookup diskstats for relevant devices
	diskstats := util.GetProcDiskstats()
	combinedDiskstat := util.ProcDiskstat{}
	/*
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
				}
			}
		} else {
	*/
	// consider all available devices
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
	}
	// }

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
}

func diskHostCollect(host *models.Host) {

}

func diskPrintHost(host *models.Host) []string {
	diskDeviceReads := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_reads", true)
	diskDeviceReadsmerged := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_readsmerged", true)
	diskDeviceSectorsread := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_sectorsread", true)
	diskDeviceTimereading := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_timereading", true)
	diskDeviceWrites := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_writes", true)
	diskDeviceWritesmerged := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_writesmerged", true)
	diskDeviceSectorswritten := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_sectorswritten", true)
	diskDeviceTimewriting := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_timewriting", true)
	diskDeviceCurrentops := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_currentops", true)
	diskDeviceTimeforops := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_timeforops", true)
	diskDeviceWeightedtimeforops := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_weightedtimeforops", true)

	result := append([]string{diskDeviceReads}, diskDeviceWrites)
	if config.Options.Verbose {
		result = append(result, diskDeviceReadsmerged, diskDeviceSectorsread, diskDeviceTimereading)
		result = append(result, diskDeviceWritesmerged, diskDeviceSectorswritten, diskDeviceTimewriting, diskDeviceCurrentops, diskDeviceTimeforops)
		result = append(result, diskDeviceWeightedtimeforops)
	}

	return result
}
