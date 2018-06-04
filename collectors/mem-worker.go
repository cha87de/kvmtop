package collectors

import (
	"github.com/cha87de/kvmtop/models"
	"github.com/cha87de/kvmtop/util"
	libvirt "github.com/libvirt/libvirt-go"
)

const pagesize = 4096

func memLookup(domain *models.Domain, libvirtDomain libvirt.Domain) {
	memStats, err := libvirtDomain.MemoryStats(uint32(libvirt.DOMAIN_MEMORY_STAT_NR), 0)
	if err != nil {
		return
	}
	var total, unused, used uint64
	for _, stat := range memStats {
		if stat.Tag == int32(libvirt.DOMAIN_MEMORY_STAT_UNUSED) {
			unused = stat.Val
		}
		if stat.Tag == int32(libvirt.DOMAIN_MEMORY_STAT_AVAILABLE) {
			total = stat.Val
		}
	}
	used = total - unused
	newMeasurementTotal := models.CreateMeasurement(total)
	domain.AddMetricMeasurement("ram_total", newMeasurementTotal)
	newMeasurementUsed := models.CreateMeasurement(used)
	domain.AddMetricMeasurement("ram_used", newMeasurementUsed)

}

func memCollect(domain *models.Domain) {
	pid := domain.PID
	stats := util.GetProcStat(pid)
	// fmt.Printf("vsize: %d, rss: %d\n", stats.VSize/1024/1024, stats.RSS*4096/1024/1024)
	// fmt.Printf("MinFlt: %d, CMinFlt: %d, MajFlt: %d, CMajFlt: %d\n", stats.MinFlt, stats.CMinFlt, stats.MajFlt, stats.CMajFlt)
	domain.AddMetricMeasurement("ram_vsize", models.CreateMeasurement(uint64(stats.VSize)))
	domain.AddMetricMeasurement("ram_rss", models.CreateMeasurement(uint64(stats.RSS*pagesize)))

	domain.AddMetricMeasurement("ram_minflt", models.CreateMeasurement(uint64(stats.MinFlt)))
	domain.AddMetricMeasurement("ram_cminflt", models.CreateMeasurement(uint64(stats.CMinFlt)))
	domain.AddMetricMeasurement("ram_majflt", models.CreateMeasurement(uint64(stats.MajFlt)))
	domain.AddMetricMeasurement("ram_cmajflt", models.CreateMeasurement(uint64(stats.CMajFlt)))
}

func memPrint(domain *models.Domain) []string {
	total := getMetricUint64(domain, "ram_total", 0)
	used := getMetricUint64(domain, "ram_used", 0)

	vsize := getMetricUint64(domain, "ram_vsize", 0)
	rss := getMetricUint64(domain, "ram_rss", 0)

	minflt := getMetricDiffUint64(domain, "ram_minflt", false)
	cminflt := getMetricDiffUint64(domain, "ram_cminflt", false)
	majflt := getMetricDiffUint64(domain, "ram_majflt", false)
	cmajflt := getMetricDiffUint64(domain, "ram_cmajflt", false)

	result := append([]string{total}, used, vsize, rss, minflt, cminflt, majflt, cmajflt)
	return result
}
