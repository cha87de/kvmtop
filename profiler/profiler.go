package profiler

import (
	"sync"
	"time"

	"github.com/cha87de/tsprofiler/impl"
	"github.com/cha87de/tsprofiler/spec"

	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
	"github.com/cha87de/kvmtop/printers"
)

var domainProfiler sync.Map

// InitializeProfiler starts the periodical profiler
func InitializeProfiler(wg *sync.WaitGroup) {
	printers.OutputOpen()

	// pull measurements in frequency
	for n := -1; config.Options.Runs == -1 || n < config.Options.Runs; n++ {
		start := time.Now()
		pickup()
		nextRun := start.Add(time.Duration(config.Options.Frequency) * time.Second)
		time.Sleep(nextRun.Sub(time.Now()))
	}

	// return from runner
	printers.OutputClose()
	wg.Done()
}

func pickup() {
	// for each domain ...
	models.Collection.Domains.Map.Range(func(key, domainRaw interface{}) bool {
		domain := domainRaw.(models.Domain)
		uuid := key.(string)

		// write measurements to profiler
		profilerRaw, found := domainProfiler.Load(uuid)
		var profiler spec.TSProfiler
		if found {
			profiler = profilerRaw.(spec.TSProfiler)
		} else {
			profiler = impl.NewProfiler(spec.Settings{
				Name:           uuid,
				BufferSize:     config.Options.Frequency * 10,
				States:         4,
				OutputFreq:     time.Duration(20) * time.Second,
				OutputCallback: profileOutput,
			})
		}

		metrics := make([]spec.TSDataMetric, 0)
		models.Collection.Collectors.Map.Range(func(nameRaw interface{}, collectorRaw interface{}) bool {
			name := nameRaw.(string)
			var util int
			if name == "cpu" {
				util = pickupCPU(domain)
			} else if name == "io" {
				util = pickupIO(domain)
			} else if name == "net" {
				util = pickupNet(domain)
			}
			metrics = append(metrics, spec.TSDataMetric{
				Name:  name,
				Value: float64(util),
			})
			return true
		})

		tsdata := spec.TSData{
			Metrics: metrics,
		}
		profiler.Put(tsdata)

		domainProfiler.Store(uuid, profiler)
		return true
	})

	// TODO clean up: remove old domains
}
