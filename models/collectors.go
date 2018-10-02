package models

import "sync"

// Collectors holds the configured collectors in a sync.Map
type Collectors struct {
	Map sync.Map
}

// Load returns the value stored as "key" identifier
func (collectors *Collectors) Load(key interface{}) (Collector, bool) {
	rawValue, ok := collectors.Map.Load(key)
	if !ok {
		return nil, ok
	}
	value := rawValue.(Collector)
	return value, ok
}

// Store writes the key value pair to the sync map
func (collectors *Collectors) Store(key string, value Collector) {
	collectors.Map.Store(key, value)
}

// Length returns the number of elements stored in the map
func (collectors *Collectors) Length() int {
	counter := 0
	collectors.Map.Range(func(_, _ interface{}) bool {
		counter++
		return true
	})
	return counter
}
