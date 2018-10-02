package iocollector

import (
	"github.com/cha87de/kvmtop/collectors"
	"github.com/cha87de/kvmtop/models"
	"github.com/cha87de/kvmtop/util"
	libvirt "github.com/libvirt/libvirt-go"
)

func ioLookup(domain *models.Domain, libvirtDomain libvirt.Domain) {
	// nothing to do
}

func ioCollect(domain *models.Domain) {
	stats := util.GetProcIO(domain.PID)
	domain.AddMetricMeasurement("io_rchar", models.CreateMeasurement(uint64(stats.Rchar)))
	domain.AddMetricMeasurement("io_wchar", models.CreateMeasurement(uint64(stats.Wchar)))
	domain.AddMetricMeasurement("io_syscr", models.CreateMeasurement(uint64(stats.Syscr)))
	domain.AddMetricMeasurement("io_syscw", models.CreateMeasurement(uint64(stats.Syscw)))
	domain.AddMetricMeasurement("io_read_bytes", models.CreateMeasurement(uint64(stats.Read_bytes)))
	domain.AddMetricMeasurement("io_write_bytes", models.CreateMeasurement(uint64(stats.Write_bytes)))
	domain.AddMetricMeasurement("io_cancelled_write_bytes", models.CreateMeasurement(uint64(stats.Cancelled_write_bytes)))
}

func ioPrint(domain *models.Domain) []string {
	rchar := collectors.GetMetricDiffUint64(domain.Measurable, "io_rchar", true)
	wchar := collectors.GetMetricDiffUint64(domain.Measurable, "io_wchar", true)
	syscr := collectors.GetMetricDiffUint64(domain.Measurable, "io_syscr", true)
	syscw := collectors.GetMetricDiffUint64(domain.Measurable, "io_syscw", true)
	readBytes := collectors.GetMetricDiffUint64(domain.Measurable, "io_read_bytes", true)
	writeBytes := collectors.GetMetricDiffUint64(domain.Measurable, "io_write_bytes", true)
	cancelledWriteBytes := collectors.GetMetricDiffUint64(domain.Measurable, "io_cancelled_write_bytes", true)

	result := append([]string{rchar}, wchar, syscr, syscw, readBytes, writeBytes, cancelledWriteBytes)
	return result
}
