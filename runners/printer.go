package runners

import (
	"bytes"
	"encoding/gob"
	"sync"
	"time"

	"fmt"

	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
)

func initializePrinter(wg *sync.WaitGroup) {
	for n := -1; config.Options.Runs == -1 || n < config.Options.Runs; n++ {
		start := time.Now()
		print()
		nextRun := start.Add(time.Duration(config.Options.Frequency) * time.Second)
		time.Sleep(nextRun.Sub(time.Now()))
	}
	wg.Done()
}

func print() {
	for _, domain := range models.Collection.Domains {
		var cores int
		if metric, ok := domain.Metrics["cpu_cores"]; ok {
			if len(metric.Values) > 0 {
				byteValue := metric.Values[0].Value
				reader := bytes.NewReader(byteValue)
				dec := gob.NewDecoder(reader)
				dec.Decode(&cores)
			}
		}

		var threadIDs []int
		if metric, ok := domain.Metrics["cpu_threadIDs"]; ok {
			if len(metric.Values) > 0 {
				byteValue := metric.Values[0].Value
				reader := bytes.NewReader(byteValue)
				dec := gob.NewDecoder(reader)
				dec.Decode(&threadIDs)
			}
		}

		fmt.Printf("%s \t %s \t %d \t %v\n", domain.UUID, domain.Name, cores, threadIDs)
	}

}
