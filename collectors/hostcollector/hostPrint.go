package hostcollector

import (
	"github.com/cha87de/kvmtop/collectors"
	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
)

func hostPrint(host *models.Host) []string {
	hostname := collectors.GetMetricString(host.Measurable, "host_name", 0)
	result := []string{hostname}

	if config.Options.Verbose {
		hostUUID := collectors.GetMetricString(host.Measurable, "host_uuid", 0)
		result = append(result, hostUUID)
	}
	return result
}
