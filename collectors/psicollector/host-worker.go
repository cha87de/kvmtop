package psicollector

import (
	"fmt"

	"kvmtop/config"
	"kvmtop/util"

	"kvmtop/models"
)

func hostLookup(host *models.Host) {
	// nothing to do
}

func hostCollect(host *models.Host) {
	hostCollectResource(host, util.ProcPressureResourceCPU, false)
	hostCollectResource(host, util.ProcPressureResourceIO, true)
	hostCollectResource(host, util.ProcPressureResourceMemory, true)
}

func hostCollectResource(host *models.Host, resource util.ProcPressureResource, queryFullMetric bool) {
	pressures := util.GetProcPressure(resource)
	metrics := []util.ProcPressureMetric{util.ProcPressureMetricSome}
	if queryFullMetric {
		metrics = append(metrics, util.ProcPressureMetricFull)
	}
	for _, metric := range metrics {
		fieldPrefix := fmt.Sprintf("psi_%s_%s", metric, resource)
		// find metric in pressureValues
		var pressure util.ProcPressure
		for _, p := range pressures {
			if p.Metric == metric {
				pressure = p
				break
			}
		}
		// copy values from pressure
		host.AddMetricMeasurement(fmt.Sprintf("%s_avg10", fieldPrefix),
			models.CreateMeasurement(pressure.Avg10))
		host.AddMetricMeasurement(fmt.Sprintf("%s_avg60", fieldPrefix),
			models.CreateMeasurement(pressure.Avg60))
		host.AddMetricMeasurement(fmt.Sprintf("%s_avg300", fieldPrefix),
			models.CreateMeasurement(pressure.Avg300))
		host.AddMetricMeasurement(fmt.Sprintf("%s_total", fieldPrefix),
			models.CreateMeasurement(pressure.Total))
	}
}

func printHost(host *models.Host) []string {
	result := []string{}
	cpu := printHostResource(host, util.ProcPressureResourceCPU, false, config.Options.Verbose)
	result = append(result, cpu...)
	io := printHostResource(host, util.ProcPressureResourceIO, true, config.Options.Verbose)
	result = append(result, io...)
	mem := printHostResource(host, util.ProcPressureResourceMemory, true, config.Options.Verbose)
	result = append(result, mem...)
	return result
}

func printHostResource(host *models.Host, resource util.ProcPressureResource, queryFullMetric bool, verbose bool) []string {
	result := []string{}
	metrics := []util.ProcPressureMetric{util.ProcPressureMetricSome}
	if queryFullMetric && verbose {
		metrics = append(metrics, util.ProcPressureMetricFull)
	}
	for _, metric := range metrics {
		fieldPrefix := fmt.Sprintf("psi_%s_%s", metric, resource)

		result = append(result, host.GetMetricFloat64(fmt.Sprintf("%s_avg60", fieldPrefix), 0))
		if verbose {
			result = append(result, host.GetMetricFloat64(fmt.Sprintf("%s_avg10", fieldPrefix), 0))
			result = append(result, host.GetMetricFloat64(fmt.Sprintf("%s_avg300", fieldPrefix), 0))
			val, _ := host.GetMetricUint64(fmt.Sprintf("%s_total", fieldPrefix), 0)
			result = append(result, val)
		}
	}
	return result
}
