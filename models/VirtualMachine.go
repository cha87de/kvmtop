package models

import (
	"bytes"
	"strings"
	"fmt"
	"log"
	"os/exec"
	"time"
	"io/ioutil"
)

// BUILDER FUNCTION
func CreateVM(ppid string) VirtualMachine {
	vm := VirtualMachine{}
	vm.ppid = ppid
	return vm
}


// CLASS VIRTUAL MACHINE
type VirtualMachine struct {
	ppid string
	libvirtname string
	vCpuTasks []string	
	nextDetailsLookup time.Time
}

func (vm *VirtualMachine) setVmName() {
    file := fmt.Sprint("/proc/", vm.ppid, "/cmdline")
    content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	str := strings.Replace(string(content), "\x00", " ", -1)
	str_arr := strings.Split(str, " ")
	vm.libvirtname = str_arr[2]
}

func (vm *VirtualMachine) Name() string {
    return vm.libvirtname
}

func (vm *VirtualMachine) Ppid() string {
    return vm.ppid
}

func (vm *VirtualMachine) VCpuTasks() []string {
    return vm.vCpuTasks
}


func (vm *VirtualMachine) setVCpus() {
	cmd := exec.Command("virsh", "--connect=qemu:///system", "qemu-monitor-command", "--hmp", vm.libvirtname, "info cpus")
	var cpuinfo bytes.Buffer
	cmd.Stdout = &cpuinfo
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	cpuinfoArr := strings.Split(cpuinfo.String(), "\n")
	var vCpuTasks []string
	for _, line := range cpuinfoArr {
		if(line == ""){
			continue
		}
		lineArr := strings.Split(line, "=")
		cpuTaskId := strings.Replace(lineArr[2], "\r", "", -1)
		cpuTaskId = strings.Replace(cpuTaskId, " ", "", -1)
		vCpuTasks = append(vCpuTasks, cpuTaskId)
	}
	vm.vCpuTasks = vCpuTasks
}

func (vm *VirtualMachine) CollectDetails(forceLookup bool) {
	// only lookup if forced or time of details is up
	if forceLookup || time.Now().After(vm.nextDetailsLookup) {
		// do lookups!			
		vm.setVmName()
		vm.setVCpus() // TODO move this to cpu collector method CollectDetails 
		
		// look through collectors
		for _, collector := range GetCollectors() {
			collector.CollectDetails(*vm)
		}
		
		// set next lookup in 10 minutes
		vm.nextDetailsLookup = time.Now().Add(time.Duration(10)*time.Minute)
	}
}

