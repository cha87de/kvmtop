package runners

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/connector"
	"github.com/cha87de/kvmtop/models"
	"github.com/cha87de/kvmtop/util"
	libvirt "github.com/libvirt/libvirt-go"
)

var processes []int

func initializeLookup(wg *sync.WaitGroup) {
	for n := -1; config.Options.Runs == -1 || n < config.Options.Runs; n++ {
		start := time.Now()
		lookup()
		nextRun := start.Add(time.Duration(config.Options.Frequency) * time.Second)
		time.Sleep(nextRun.Sub(time.Now()))
	}
	wg.Done()
}

func lookup() {
	// initialize models
	if models.Collection.Domains == nil {
		models.Collection.Domains = make(map[string]*models.Domain)
	}

	// query libvirt
	doms, err := connector.Libvirt.Connection.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	libvirtDomains := make(map[string]libvirt.Domain)
	if err != nil {
		log.Printf("Cannot get list of domains form libvirt.")
		return
	}

	// create list of cached domains
	domIDs := make([]string, 0, len(models.Collection.Domains))
	for id := range models.Collection.Domains {
		domIDs = append(domIDs, id)
	}

	// update process list
	processes = util.GetProcessList()

	// update domain list
	for _, dom := range doms {
		domain, err := handleDomain(dom)
		libvirtDomains[domain.UUID] = dom
		if err != nil {
			continue
		}
		domIDs = removeFromArray(domIDs, domain.UUID)
	}

	// remove cached but not existent domains
	for _, id := range domIDs {
		delete(models.Collection.Domains, id)
	}

	// call collector lookup functions
	for _, collector := range models.Collection.Collectors {
		collector.Lookup(models.Collection.Domains, libvirtDomains)
	}

}

func handleDomain(dom libvirt.Domain) (*models.Domain, error) {
	uuid, err := dom.GetUUIDString()
	if err != nil {
		return nil, err
	}

	name, err := dom.GetName()
	if err != nil {
		return nil, err
	}

	if domain, ok := models.Collection.Domains[uuid]; ok {
		domain.Name = name
		models.Collection.Domains[uuid] = domain
	} else {
		models.Collection.Domains[uuid] = &models.Domain{
			UUID: string(uuid),
			Name: name,
		}
	}

	// lookup PID
	var pid int
	for _, process := range processes {
		cmdline := util.GetCmdLine(process)
		if cmdline != "" && strings.Contains(cmdline, name) {
			// fmt.Printf("Found PID %d for instance %s (cmdline: %s)", process, name, cmdline)
			pid = process
			break
		}
	}
	models.Collection.Domains[uuid].PID = pid

	return models.Collection.Domains[uuid], nil
}

func removeFromArray(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
