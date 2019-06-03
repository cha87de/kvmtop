package models

import "sync"

// NewCollectors instantiates a new Collectors list
func NewCollectors() *Collectors {
	return &Collectors{
		access:     sync.Mutex{},
		collectors: make(map[interface{}]Collector),
	}
}

// Collectors holds the configured collectors in a Map
type Collectors struct {
	// Map sync.Map
	collectors map[interface{}]Collector
	access     sync.Mutex
}

// Load returns the value stored as "key" identifier
func (collectors *Collectors) Load(key interface{}) (Collector, bool) {
	collectors.access.Lock()
	defer collectors.access.Unlock()
	Collector, exists := collectors.collectors[key]
	if !exists {
		return nil, false
	}
	return Collector, true
}

// Store writes the key value pair to the sync map
func (collectors *Collectors) Store(key string, value Collector) {
	collectors.access.Lock()
	defer collectors.access.Unlock()
	//collectors.Map.Store(key, value)
	collectors.collectors[key] = value
}

// Length returns the number of elements stored in the map
func (collectors *Collectors) Length() int {
	return len(collectors.collectors)
}

// Range loops over the items and executes the given function
func (collectors *Collectors) Range(f func(key interface{}, value Collector) bool) {
	collectors.access.Lock()
	defer collectors.access.Unlock()
	for key, collector := range collectors.collectors {
		f(key, collector)
	}
}

// Delete removes the element given by ke from the collectors list
func (collectors *Collectors) Delete(key string) {
	collectors.access.Lock()
	defer collectors.access.Unlock()
	delete(collectors.collectors, key)
}
