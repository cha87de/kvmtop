package models

import (
	"sync"

	libvirt "github.com/libvirt/libvirt-go"
)

// LibvirtDomains holds the collected libvirt domains in a sync.Map
type LibvirtDomains struct {
	Map sync.Map
}

// Load returns the value stored as "key" identifier
func (libvirtDomains *LibvirtDomains) Load(key interface{}) (libvirt.Domain, bool) {
	rawValue, ok := libvirtDomains.Map.Load(key)
	if !ok {
		return libvirt.Domain{}, ok
	}
	value := rawValue.(libvirt.Domain)
	return value, ok
}

// Store writes the key value pair to the sync map
func (libvirtDomains *LibvirtDomains) Store(key string, value libvirt.Domain) {
	libvirtDomains.Map.Store(key, value)
}

// Length returns the number of elements stored in the map
func (libvirtDomains *LibvirtDomains) Length() int {
	counter := 0
	libvirtDomains.Map.Range(func(_, _ interface{}) bool {
		counter++
		return true
	})
	return counter
}
