package profiler

import (
	"strconv"
	"sync"
	"time"

	"github.com/cha87de/tsprofiler/impl"
	"github.com/cha87de/tsprofiler/spec"

	"github.com/cha87de/kvmtop/collectors/cpucollector"
	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
)

var domainProfiler sync.Map

// InitializeProfiler starts the periodical profiler
func InitializeProfiler(wg *sync.WaitGroup) {

	// pull measurements in frequency
	for n := -1; config.Options.Runs == -1 || n < config.Options.Runs; n++ {
		start := time.Now()
		pickup()
		nextRun := start.Add(time.Duration(config.Options.Frequency) * time.Second)
		time.Sleep(nextRun.Sub(time.Now()))
	}

	// return from runner
	wg.Done()
}

func pickup() {
	// for each domain ...
	models.Collection.Domains.Map.Range(func(key, domainRaw interface{}) bool {
		domain := domainRaw.(models.Domain)
		uuid := key.(string)

		// get CPU util
		cputimeAllCores, _ := strconv.Atoi(cpucollector.CpuPrintThreadMetric(&domain, "cpu_times"))
		queuetimeAllCores, _ := strconv.Atoi(cpucollector.CpuPrintThreadMetric(&domain, "cpu_runqueues"))
		cpuUtil := cputimeAllCores + queuetimeAllCores

		// write measurements in a ring
		profilerRaw, found := domainProfiler.Load(uuid)
		var profiler spec.TSProfiler
		if found {
			profiler = profilerRaw.(spec.TSProfiler)
		} else {
			profiler = impl.NewSimpleProfiler(spec.Settings{
				Frequency: config.Options.Frequency * 10,
				Name:      uuid + "_cpu",
			})
		}

		profiler.Put(spec.TSData{
			Value: float64(cpuUtil),
		})

		domainProfiler.Store(uuid, profiler)
		return true
	})

	// TODO clean up: remove old domains
}
