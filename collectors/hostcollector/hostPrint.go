package hostcollector

import (
	"kvmtop/config"
	"kvmtop/models"
)

func hostPrint(host *models.Host) []string {
	hostname := host.GetMetricString("host_name", 0)
	result := []string{hostname}

	if config.Options.Verbose {
		hostUUID := host.GetMetricString("host_uuid", 0)
		result = append(result, hostUUID)
	}
	return result
}
