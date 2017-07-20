package runners

import (
	"sync"
	"time"

	"github.com/cha87de/kvmtop/config"
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
	// todo
}
