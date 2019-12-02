package memcollector

import (
	"kvmtop/config"
	"kvmtop/models"
)

func domainPrint(domain *models.Domain) []string {
	total, _ := domain.GetMetricUint64("ram_total", 0)
	used, _ := domain.GetMetricUint64("ram_used", 0)

	vsize, _ := domain.GetMetricUint64("ram_vsize", 0)
	rss, _ := domain.GetMetricUint64("ram_rss", 0)

	minflt := domain.GetMetricDiffUint64("ram_minflt", false)
	cminflt := domain.GetMetricDiffUint64("ram_cminflt", false)
	majflt := domain.GetMetricDiffUint64("ram_majflt", false)
	cmajflt := domain.GetMetricDiffUint64("ram_cmajflt", false)

	result := append([]string{total}, used)
	if config.Options.Verbose {
		result = append(result, vsize, rss, minflt, cminflt, majflt, cmajflt)
	}

	return result
}
