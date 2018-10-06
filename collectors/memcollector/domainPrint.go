package memcollector

import (
	"github.com/cha87de/kvmtop/collectors"
	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
)

func domainPrint(domain *models.Domain) []string {
	total := collectors.GetMetricUint64(domain.Measurable, "ram_total", 0)
	used := collectors.GetMetricUint64(domain.Measurable, "ram_used", 0)

	vsize := collectors.GetMetricUint64(domain.Measurable, "ram_vsize", 0)
	rss := collectors.GetMetricUint64(domain.Measurable, "ram_rss", 0)

	minflt := collectors.GetMetricDiffUint64(domain.Measurable, "ram_minflt", false)
	cminflt := collectors.GetMetricDiffUint64(domain.Measurable, "ram_cminflt", false)
	majflt := collectors.GetMetricDiffUint64(domain.Measurable, "ram_majflt", false)
	cmajflt := collectors.GetMetricDiffUint64(domain.Measurable, "ram_cmajflt", false)

	result := append([]string{total}, used)
	if config.Options.Verbose {
		result = append(result, vsize, rss, minflt, cminflt, majflt, cmajflt)
	}

	return result
}
