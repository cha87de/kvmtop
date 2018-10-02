package models

import "sync"

// Domains holds the collected Domains in a sync.Map
type Domains struct {
	Map sync.Map
}

// Load returns the value stored as "key" identifier
func (domains *Domains) Load(key interface{}) (Domain, bool) {
	rawValue, ok := domains.Map.Load(key)
	if !ok {
		return Domain{}, ok
	}
	value := rawValue.(Domain)
	return value, ok
}

// Store writes the key value pair to the sync map
func (domains *Domains) Store(key string, value Domain) {
	domains.Map.Store(key, value)
}

// Length returns the number of elements stored in the map
func (domains *Domains) Length() int {
	counter := 0
	domains.Map.Range(func(_, _ interface{}) bool {
		counter++
		return true
	})
	return counter
}
