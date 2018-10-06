package profiler

import (
	"fmt"
	"sync"
	"time"

	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
)

// InitializeProfiler starts the periodical profiler
func InitializeProfiler(wg *sync.WaitGroup) {
	// start continuously printing values
	for n := -1; config.Options.Runs == -1 || n < config.Options.Runs; n++ {
		start := time.Now()
		profile()
		nextRun := start.Add(time.Duration(config.Options.Frequency*10) * time.Second)
		time.Sleep(nextRun.Sub(time.Now()))
	}

	// return from runner
	wg.Done()
}

func profile() {

	// TODO implement profiler here

	// for each domain ...
	models.Collection.Domains.Map.Range(func(key, domainRaw interface{}) bool {
		domain := domainRaw.(models.Domain)
		fmt.Printf("%+v\n", domain)
		return true
	})

}
