package hostcollector

import (
	"github.com/cha87de/kvmtop/models"
)

func domainPrint(domain *models.Domain) []string {
	host := domain.GetMetricString("host_name", 0)
	result := append([]string{host})
	return result
}
