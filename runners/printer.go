package runners

import (
	"sync"
	"time"

	"fmt"

	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
)

func initializePrinter(wg sync.WaitGroup) {
	for n := -1; config.Options.Runs == -1 || n < config.Options.Runs; n++ {
		start := time.Now()
		print()
		nextRun := start.Add(time.Duration(config.Options.Frequency) * time.Second)
		time.Sleep(nextRun.Sub(time.Now()))
	}
	wg.Done()
}

func print() {
	fmt.Printf("Domains: %+v\n", models.Collection.Domains)
}
