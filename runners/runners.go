package runners

import (
	"sync"
)

// InitializeRunners starts necessary runners as threads
func InitializeRunners() {

	var wg sync.WaitGroup
	wg.Add(1) // terminate when first thread terminates

	go initializeLookup(&wg)
	go initializeCollect(&wg)
	go initializePrinter(&wg)

	wg.Wait()
}
