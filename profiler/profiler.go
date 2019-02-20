package profiler

import (
	"sync"
	"time"

	tsprofilerApi "github.com/cha87de/tsprofiler/api"
	tsprofilerModels "github.com/cha87de/tsprofiler/models"
	tsprofiler "github.com/cha87de/tsprofiler/profiler"

	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
	"github.com/cha87de/kvmtop/printers"
	"github.com/cha87de/kvmtop/util"
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

	// create list of cached profilers
	domIDs := make([]string, 0)
	domainProfiler.Range(func(key, _ interface{}) bool {
		domIDs = append(domIDs, key.(string))
		return true
	})

	// for each domain ...
	models.Collection.Domains.Map.Range(func(key, domainRaw interface{}) bool {
		domain := domainRaw.(models.Domain)
		uuid := key.(string)

		// get or create profiler
		profilerRaw, found := domainProfiler.Load(uuid)
		var profiler tsprofilerApi.TSProfiler
		if found {
			profiler = profilerRaw.(tsprofilerApi.TSProfiler)
		} else {
			profiler = tsprofiler.NewProfiler(tsprofilerModels.Settings{
				Name:           uuid,
				BufferSize:     config.Options.Profiler.BufferSize, // default: 10, with default 1s frequency => every 10s
				States:         config.Options.Profiler.States,     // default: 4
				History:        config.Options.Profiler.History,    // default: 1
				FilterStdDevs:  config.Options.Profiler.FilterStdDevs,
				FixBound:       config.Options.Profiler.FixedBound,
				OutputFreq:     config.Options.Profiler.OutputFreq,
				OutputCallback: profileOutput,
				PeriodSize:     []int{6, 60}, // 1min, 1h
			})
		}

		// pick up collector measurement
		metrics := make([]tsprofilerModels.TSInputMetric, 0)
		models.Collection.Collectors.Map.Range(func(nameRaw interface{}, collectorRaw interface{}) bool {
			name := nameRaw.(string)
			var util, min, max int
			if name == "cpu" {
				util, min, max = pickupCPU(domain)
			} else if name == "io" {
				util, min, max = pickupIO(domain)
			} else if name == "net" {
				util, min, max = pickupNet(domain)
			}

			metrics = append(metrics, tsprofilerModels.TSInputMetric{
				Name:     name,
				Value:    float64(util),
				FixedMin: float64(min),
				FixedMax: float64(max),
			})
			return true
		})

		// send measurement to profiler
		tsdata := tsprofilerModels.TSInput{
			Metrics: metrics,
		}
		profiler.Put(tsdata)

		// store profiler
		domainProfiler.Store(uuid, profiler)

		// mark domain as considered by removing from cache
		domIDs = util.RemoveFromArray(domIDs, uuid)

		return true
	})

	// remove cached profilers for not existent domains
	for _, uuid := range domIDs {
		profilerRaw, found := domainProfiler.Load(uuid)
		if found {
			profiler := profilerRaw.(tsprofilerApi.TSProfiler)
			profiler.Terminate()
		}
		domainProfiler.Delete(uuid)
	}
}
