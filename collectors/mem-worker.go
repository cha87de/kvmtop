package collectors

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/cha87de/kvmtop/models"
	"github.com/cha87de/kvmtop/util"
	libvirt "github.com/libvirt/libvirt-go"
)

var PAGESIZE = 4096

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
	domain.AddMetricMeasurement("ram_rss", models.CreateMeasurement(uint64(stats.RSS*PAGESIZE)))
	domain.AddMetricMeasurement("ram_minflt", models.CreateMeasurement(uint64(stats.MinFlt)))
	domain.AddMetricMeasurement("ram_cminflt", models.CreateMeasurement(uint64(stats.CMinFlt)))
	domain.AddMetricMeasurement("ram_majflt", models.CreateMeasurement(uint64(stats.MajFlt)))
	domain.AddMetricMeasurement("ram_cmajflt", models.CreateMeasurement(uint64(stats.CMajFlt)))
}

func memPrint(domain *models.Domain) []string {
	total := memPrintMetric(domain, "ram_total")
	used := memPrintMetric(domain, "ram_used")

	vsize := memPrintMetric(domain, "ram_vsize")
	rss := memPrintMetric(domain, "ram_rss")
	minflt := memPrintMetric(domain, "ram_minflt")
	cminflt := memPrintMetric(domain, "ram_cminflt")
	majflt := memPrintMetric(domain, "ram_majflt")
	cmajflt := memPrintMetric(domain, "ram_cmajflt")

	result := append([]string{total}, used, vsize, rss, minflt, cminflt, majflt, cmajflt)
	return result
}

func memPrintMetric(domain *models.Domain, metric string) string {
	var output string
	if metric, ok := domain.GetMetric(metric); ok {
		if len(metric.Values) > 0 {
			byteValue := metric.Values[0].Value
			reader := bytes.NewReader(byteValue)
			decoder := gob.NewDecoder(reader)
			var value uint64
			decoder.Decode(&value)
			output = fmt.Sprintf("%d", value/1024)
		}
	}
	return output
}
