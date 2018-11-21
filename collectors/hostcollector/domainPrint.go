package hostcollector

import (
	"github.com/cha87de/kvmtop/collectors"
	"github.com/cha87de/kvmtop/models"
)

func domainPrint(domain *models.Domain) []string {
	host := collectors.GetMetricString(domain.Measurable, "host_name", 0)
	result := append([]string{host})
	return result
}
