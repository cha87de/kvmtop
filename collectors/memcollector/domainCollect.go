package memcollector

import (
	"github.com/cha87de/kvmtop/models"
	"github.com/cha87de/kvmtop/util"
)

func domainCollect(domain *models.Domain) {
	pid := domain.PID
	stats := util.GetProcPIDStat(pid)
	// fmt.Printf("vsize: %d, rss: %d\n", stats.VSize/1024/1024, stats.RSS*4096/1024/1024)
	// fmt.Printf("MinFlt: %d, CMinFlt: %d, MajFlt: %d, CMajFlt: %d\n", stats.MinFlt, stats.CMinFlt, stats.MajFlt, stats.CMajFlt)
	domain.AddMetricMeasurement("ram_vsize", models.CreateMeasurement(uint64(stats.VSize)))
	domain.AddMetricMeasurement("ram_rss", models.CreateMeasurement(uint64(stats.RSS*pagesize)))

	domain.AddMetricMeasurement("ram_minflt", models.CreateMeasurement(uint64(stats.MinFlt)))
	domain.AddMetricMeasurement("ram_cminflt", models.CreateMeasurement(uint64(stats.CMinFlt)))
	domain.AddMetricMeasurement("ram_majflt", models.CreateMeasurement(uint64(stats.MajFlt)))
	domain.AddMetricMeasurement("ram_cmajflt", models.CreateMeasurement(uint64(stats.CMajFlt)))
}
