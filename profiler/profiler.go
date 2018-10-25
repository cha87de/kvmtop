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

		cpuUtil := pickupCPU(domain)
		ioUtil := pickupIO(domain)
		netUtil := pickupNet(domain)

		// write measurements to profiler
		profilerRaw, found := domainProfiler.Load(uuid)
		var profiler spec.TSProfiler
		if found {
			profiler = profilerRaw.(spec.TSProfiler)
		} else {
			profiler = impl.NewSimpleProfiler(spec.Settings{
				Name:           uuid,
				BufferSize:     config.Options.Frequency * 10,
				OutputFreq:     time.Duration(1) * time.Minute,
				OutputCallback: profileOutput,
			})
		}

		profiler.Put(spec.TSData{
			CPU: float64(cpuUtil),
			IO:  float64(ioUtil),
			Net: float64(netUtil),
		})

		domainProfiler.Store(uuid, profiler)
		return true
	})

	// TODO clean up: remove old domains
}
