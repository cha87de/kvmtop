package runners

import (
	"sync"
)

// InitializeRunners starts necessary runners as threads
func InitializeRunners() {

	var wg sync.WaitGroup
	wg.Add(1) // terminate when first thread terminates

	go InitializeLookup(&wg)
	go InitializeCollect(&wg)
	go InitializePrinter(&wg)

	wg.Wait()
}
