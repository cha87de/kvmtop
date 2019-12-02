package iocollector

import (
	"kvmtop/config"
	"kvmtop/models"
	"kvmtop/util"
	libvirt "github.com/libvirt/libvirt-go"
)

func ioLookup(domain *models.Domain, libvirtDomain libvirt.Domain) {
	// nothing to do
}

func ioCollect(domain *models.Domain) {
	stats := util.GetProcPIDIO(domain.PID)
	domain.AddMetricMeasurement("io_rchar", models.CreateMeasurement(uint64(stats.Rchar)))
	domain.AddMetricMeasurement("io_wchar", models.CreateMeasurement(uint64(stats.Wchar)))
	domain.AddMetricMeasurement("io_syscr", models.CreateMeasurement(uint64(stats.Syscr)))
	domain.AddMetricMeasurement("io_syscw", models.CreateMeasurement(uint64(stats.Syscw)))
	domain.AddMetricMeasurement("io_read_bytes", models.CreateMeasurement(uint64(stats.Read_bytes)))
	domain.AddMetricMeasurement("io_write_bytes", models.CreateMeasurement(uint64(stats.Write_bytes)))
	domain.AddMetricMeasurement("io_cancelled_write_bytes", models.CreateMeasurement(uint64(stats.Cancelled_write_bytes)))
}

func ioPrint(domain *models.Domain) []string {
	rchar := domain.GetMetricDiffUint64("io_rchar", true)
	wchar := domain.GetMetricDiffUint64("io_wchar", true)
	syscr := domain.GetMetricDiffUint64("io_syscr", true)
	syscw := domain.GetMetricDiffUint64("io_syscw", true)
	readBytes := domain.GetMetricDiffUint64("io_read_bytes", true)
	writeBytes := domain.GetMetricDiffUint64("io_write_bytes", true)
	cancelledWriteBytes := domain.GetMetricDiffUint64("io_cancelled_write_bytes", true)

	result := append([]string{readBytes}, writeBytes)
	if config.Options.Verbose {
		result = append(result, rchar, wchar, syscr, syscw, cancelledWriteBytes)
	}
	return result
}
