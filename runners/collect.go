package runners

import (
	"sync"

	"kvmtop/models"
)

// InitializeCollect starts the periodic collect calls
func InitializeCollect(wg *sync.WaitGroup) {
	for {
		// wait with execution for lookup routine
		_, ok := <-initialLookupDone
		if !ok {
			wg.Done()
			return
		}
		Collect()
	}
}

// Collect runs one collect cycle to measure frequently changing metrics
func Collect() {
	// initialize models
	/*if models.Collection.Domains.Length() <= 0 {
		// wait for lookup to create domains
		return
	}*/

	// run collectors
	models.Collection.Collectors.Range(func(_ interface{}, collector models.Collector) bool {
		go collector.Collect()
		return true
	})
}
