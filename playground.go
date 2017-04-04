package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"time"

	"os"

	libvirt "github.com/libvirt/libvirt-go"
)

type virtualMachine struct {
	name                  string
	uuid                  string
	measurementTime       time.Time
	cputimePerCore        []int64
	runqueuePerCore       []int64
	libvirtCputimePerCore []uint64
}

var frequency = 10

func main() {
	hostname, _ := os.Hostname()

	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		log.Fatal("Cannot connect to libvirt.")
	}
	defer conn.Close()

	regThreadID := regexp.MustCompile("thread_id=([0-9]*)\\s")
	regSchedStat := regexp.MustCompile("[0-9]*")

	var vms map[string]virtualMachine
	vms = make(map[string]virtualMachine)

	for {
		start := time.Now()

		doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
		if err != nil {
			log.Fatal("Cannot get list of domains form libvirt.")
		}
		//fmt.Printf("%d running domains:\n", len(doms))

		for _, dom := range doms {
			// query data from libvirt
			name, err := dom.GetName()
			if err != nil {
				continue
			}
			uuid, err := dom.GetUUIDString()
			if err != nil {
				continue
			}
			vcpus, err := dom.GetVcpus()
			if err != nil {
				continue
			}
			vCPUThreads, err := dom.QemuMonitorCommand("info cpus", libvirt.DOMAIN_QEMU_MONITOR_COMMAND_HMP)
			if err != nil {
				continue
			}
			cpuCoreThreadIds := regThreadID.FindAllStringSubmatch(vCPUThreads, -1)

			// get or create virtualMachine struct
			var vm virtualMachine
			if _vm, ok := vms[name]; ok {
				vm = _vm
			}
			vm.name = name
			vm.uuid = uuid

			// save old measurements
			cputimePerCorePrevious := make([]int64, len(vm.cputimePerCore))
			copy(cputimePerCorePrevious, vm.cputimePerCore)
			runqueuePerCorePrevious := make([]int64, len(vm.runqueuePerCore))
			copy(runqueuePerCorePrevious, vm.runqueuePerCore)
			measurementTimePrevious := vm.measurementTime
			libvirtCputimePerCorePrevious := make([]uint64, len(vm.libvirtCputimePerCore))
			copy(libvirtCputimePerCorePrevious, vm.libvirtCputimePerCore)

			// store new measurements
			for i := 0; i < len(vcpus); i++ {
				// store libvirt cputime
				if len(vm.libvirtCputimePerCore) > i {
					vm.libvirtCputimePerCore[i] = vcpus[i].CpuTime
				} else {
					vm.libvirtCputimePerCore = append(vm.libvirtCputimePerCore, vcpus[i].CpuTime)
				}

				// read and store schedstat values from /proc
				if len(cpuCoreThreadIds) > i && len(cpuCoreThreadIds[i]) >= 2 {
					threadID := cpuCoreThreadIds[i][1]
					schedstatFile := fmt.Sprint("/proc/", threadID, "/schedstat")
					schedstatFileContent, _ := ioutil.ReadFile(schedstatFile)
					schedStatCounters := regSchedStat.FindAllStringSubmatch(string(schedstatFileContent), -1)
					if len(schedStatCounters) < 2 {
						log.Fatal(fmt.Sprintf("schedstat file unreadable: %s", schedstatFileContent))
					}
					cputime, _ := strconv.ParseInt(schedStatCounters[0][0], 10, 64)
					runqueue, _ := strconv.ParseInt(schedStatCounters[1][0], 10, 64)
					if len(vm.cputimePerCore) > i {
						vm.cputimePerCore[i] = cputime
					} else {
						vm.cputimePerCore = append(vm.cputimePerCore, cputime)
					}
					if len(vm.runqueuePerCore) > i {
						vm.runqueuePerCore[i] = runqueue
					} else {
						vm.runqueuePerCore = append(vm.runqueuePerCore, runqueue)
					}
					// test output for comparison of libvirt and proc cputime value
					//fmt.Printf("\t\tcputime schedstat %d\tlibvirt %d\tdiff: %.1f\n", cputime, vcpus[i].CpuTime, float64(uint64(cputime)-vcpus[i].CpuTime)/1000000000)
				}
			}

			// calculate diffs and print output
			vm.measurementTime = time.Now()
			timeDiff := vm.measurementTime.Sub(measurementTimePrevious).Seconds()
			//fmt.Printf("timeDiff: %f", timeDiff)
			timeConversionFactor := 1000000000 / timeDiff
			if len(cputimePerCorePrevious) > 0 && len(runqueuePerCorePrevious) > 0 {
				var cputimeLibvirt float64
				for i := 0; i < len(vm.libvirtCputimePerCore); i++ {
					cputimeLibvirt = cputimeLibvirt + float64(vm.libvirtCputimePerCore[i]-libvirtCputimePerCorePrevious[i])/timeConversionFactor
				}
				var cputimeDiff, runqueueDiff, vmCputimeDiff float64
				for i := 0; i < len(vm.cputimePerCore); i++ {
					cputimeDiff = cputimeDiff + float64(vm.cputimePerCore[i]-cputimePerCorePrevious[i])/timeConversionFactor
					runqueueDiff = runqueueDiff + float64(vm.runqueuePerCore[i]-runqueuePerCorePrevious[i])/timeConversionFactor
				}
				cputimeDiff = cputimeDiff / float64(len(vm.cputimePerCore)) * 100
				runqueueDiff = runqueueDiff / float64(len(vm.cputimePerCore)) * 100
				vmCputimeDiff = cputimeDiff + runqueueDiff
				cputimeLibvirt = cputimeLibvirt / float64(len(vm.libvirtCputimePerCore)) * 100
				//fmt.Printf("cputimeDiff %.0f%% runqueueDiff %.0f%%\n", cputimeDiff, runqueueDiff)
				fmt.Printf("kvmtop,host=%s,vm=%s,level=virtual cpu_cs=%d,cpu_vm=%.0f,cpu_pm=%.0f,cpu_st=%.0f,cpu_libvirt=%.0f", hostname, vm.name, len(vcpus), vmCputimeDiff, cputimeDiff, runqueueDiff, cputimeLibvirt)
			}

			// cleanup and store vm for next cycle
			vms[name] = vm
			dom.Free()
		}

		// cycle for all vms done. sleeping ...
		nextRun := start.Add(time.Duration(frequency) * time.Second)
		time.Sleep(nextRun.Sub(time.Now()))
	}
}
