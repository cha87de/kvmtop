package runners

import (
	"sync"
	"time"

	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
)

// InitializeCollect starts the periodic collect calls
func InitializeCollect(wg *sync.WaitGroup) {
	for n := -1; config.Options.Runs == -1 || n < config.Options.Runs; n++ {
		start := time.Now()
		Collect()
		nextRun := start.Add(time.Duration(config.Options.Frequency) * time.Second)
		time.Sleep(nextRun.Sub(time.Now()))
	}
	wg.Done()
}

// Collect runs one collect cycle to measure frequently changing metrics
func Collect() {
	// initialize models
	if models.Collection.Domains.Length() <= 0 {
		// wait for lookup to create domains
		return
	}

	// run collectors
	models.Collection.Collectors.Map.Range(func(_, collectorRaw interface{}) bool {
		collector := collectorRaw.(models.Collector)
		go collector.Collect()
		return true
	})
}
