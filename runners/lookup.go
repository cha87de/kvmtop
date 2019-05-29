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
var lookupDone chan bool

// InitializeLookup starts the periodic lookup calls
func InitializeLookup(wg *sync.WaitGroup) {
	lookupDone = make(chan bool)
	for n := -1; config.Options.Runs == -1 || n < config.Options.Runs; n++ {
		// execution, then sleep
		start := time.Now()
		Lookup()
		lookupDone <- true
		freq := float32(config.Options.Frequency)
		if n <= 1 {
			// first run, half frequency only
			freq = freq / 2
		}
		nextRun := start.Add(time.Duration(freq) * time.Second)
		time.Sleep(nextRun.Sub(time.Now()))
	}
	close(lookupDone)
	wg.Done()
}

// Lookup runs one lookup cycle to detect rather static metrics
func Lookup() {
	// query libvirt
	doms, err := connector.Libvirt.Connection.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		log.Printf("Cannot get list of domains form libvirt.")
		return
	}

	// create list of cached domains
	domIDs := make([]string, 0, models.Collection.Domains.Length())

	models.Collection.Domains.Map.Range(func(key, _ interface{}) bool {
		domIDs = append(domIDs, key.(string))
		return true
	})

	// update process list
	processes = util.GetProcessList()

	// update domain list
	for _, dom := range doms {
		domain, err := handleDomain(dom)
		models.Collection.LibvirtDomains.Store(domain.UUID, dom)
		if err != nil {
			continue
		}
		domIDs = util.RemoveFromArray(domIDs, domain.UUID)
	}

	// remove cached but not existent domains
	for _, id := range domIDs {
		models.Collection.Domains.Map.Delete(id)
	}

	// call collector lookup functions
	models.Collection.Collectors.Map.Range(func(_, collectorRaw interface{}) bool {
		collector := collectorRaw.(models.Collector)
		collector.Lookup()
		return true
	})

}

func handleDomain(dom libvirt.Domain) (models.Domain, error) {
	uuid, err := dom.GetUUIDString()
	if err != nil {
		return models.Domain{}, err
	}

	name, err := dom.GetName()
	if err != nil {
		return models.Domain{}, err
	}

	// lookup or create domain
	var domain models.Domain
	var ok bool
	if domain, ok = models.Collection.Domains.Load(uuid); ok {
		domain.Name = name
	} else {
		domain = models.Domain{
			UUID:       string(uuid),
			Name:       name,
			Measurable: &models.Measurable{},
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
	domain.PID = pid

	// write back domain
	models.Collection.Domains.Store(uuid, domain)

	return domain, nil
}
