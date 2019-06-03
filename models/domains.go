package models

import "sync"

// NewDomains instantiates a new Domains list
func NewDomains() *Domains {
	return &Domains{
		access:  sync.Mutex{},
		domains: make(map[interface{}]Domain),
	}
}

// Domains holds the collected Domains in a map
type Domains struct {
	//Map sync.Map
	access  sync.Mutex
	domains map[interface{}]Domain
}

// Load returns the value stored as "key" identifier
func (domains *Domains) Load(key interface{}) (Domain, bool) {
	domains.access.Lock()
	defer domains.access.Unlock()
	domain, exists := domains.domains[key]
	if !exists {
		return Domain{}, false
	}
	return domain, true
}

// Store writes the key value pair to the sync map
func (domains *Domains) Store(key string, value Domain) {
	domains.access.Lock()
	defer domains.access.Unlock()
	//domains.Map.Store(key, value)
	domains.domains[key] = value
}

// Length returns the number of elements stored in the map
func (domains *Domains) Length() int {
	return len(domains.domains)
}

// Range loops over the items and executes the given function
func (domains *Domains) Range(f func(key, value interface{}) bool) {
	domains.access.Lock()
	defer domains.access.Unlock()
	for key, domain := range domains.domains {
		f(key, domain)
	}
}

// Delete removes the element given by ke from the domains list
func (domains *Domains) Delete(key string) {
	domains.access.Lock()
	defer domains.access.Unlock()
	delete(domains.domains, key)
}
