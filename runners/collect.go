package runners

import (
	"sync"
	"time"

	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
)

func initializeCollect(wg *sync.WaitGroup) {
	for n := -1; config.Options.Runs == -1 || n < config.Options.Runs; n++ {
		start := time.Now()
		collect()
		nextRun := start.Add(time.Duration(config.Options.Frequency) * time.Second)
		time.Sleep(nextRun.Sub(time.Now()))
	}
	wg.Done()
}

func collect() {
	// initialize models
	if models.Collection.Domains == nil {
		// wait for lookup to create domains
		return
	}

	// run domain collectors for each domain
	for _, domain := range models.Collection.Domains {
		for _, collector := range models.Collection.Collectors {
			go collector.Collect(domain)
		}
	}
}
