package diskcollector

import (
	"path/filepath"
	"strings"

	"github.com/cha87de/kvmtop/collectors"

	"github.com/cha87de/kvmtop/models"

	"github.com/cha87de/kvmtop/util"
)

func diskHostLookup(host *models.Host) {

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

	// lookup diskstats for relevant devices
	diskstats := util.GetProcDiskstats()
	combinedDiskstat := util.ProcDiskstat{}
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
	disk_device_reads := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_reads", true)
	disk_device_readsmerged := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_readsmerged", true)
	disk_device_sectorsread := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_sectorsread", true)
	disk_device_timereading := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_timereading", true)
	disk_device_writes := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_writes", true)
	disk_device_writesmerged := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_writesmerged", true)
	disk_device_sectorswritten := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_sectorswritten", true)
	disk_device_timewriting := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_timewriting", true)
	disk_device_currentops := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_currentops", true)
	disk_device_timeforops := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_timeforops", true)
	disk_device_weightedtimeforops := collectors.GetMetricDiffUint64(host.Measurable, "disk_device_weightedtimeforops", true)

	result := append([]string{disk_device_reads}, disk_device_readsmerged, disk_device_sectorsread, disk_device_timereading, disk_device_writes)
	result = append(result, disk_device_writesmerged, disk_device_sectorswritten, disk_device_timewriting, disk_device_currentops, disk_device_timeforops)
	result = append(result, disk_device_weightedtimeforops)
	return result
}
