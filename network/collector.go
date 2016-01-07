package network

import (
	"github.com/cha87de/kvmtop/models"
	"fmt"
	"bytes"
)

var (
	// store stats for calculating diffs
	vms map[string] VirtualMachineExtended
)

type NetworkCollector struct {

}

func (collector NetworkCollector) Collect(vm models.VirtualMachine) (string, error) {
		vmx := vms[vm.Name()]
		
		// new current counters
		newStats, err := vmx.lookupStats()
		if err != nil{
			return "", err
		}
	
		// get old counters
		oldStats := vmx.Statistic
		if len(oldStats) < 1{
			vmx.Statistic = newStats
			vms[vm.Name()] = vmx
			return "-", nil
		}

		// calculate diff between new and old counters
		var utilSum int64 = 0
		for itemName, newStat := range newStats {
			if oldStat, exists := oldStats[itemName]; exists{
				utilItem := newStat.DiffPerTime(oldStat)
				utilSum = utilSum + utilItem
			}
		}
		
		// set newStats as oldStats
		vmx.Statistic = newStats
		vms[vm.Name()] = vmx
		
		utilMB := (float64(utilSum)/1024/1024)
		result := fmt.Sprintf("%.2fMB/s", utilMB)
		return result, nil			

}

func (collector NetworkCollector) CollectDetails(vm models.VirtualMachine) {
	// lookup network interfaces for all virtual machines
	iflist, err := readItems(vm)
	if err != nil {
		return
	}
	if vmx, exists := vms[vm.Name()]; exists {
		vmx.Items = iflist
		vms[vm.Name()] = vmx
	}else{
		vms[vm.Name()] = VirtualMachineExtended{vm, iflist, nil}
	}

}

func DefineFlags() {
	//flag.BoolVar(&CPU_EACH, "cpu-each", CPU_EACH, "CPU each")
}

func PrintHeader(buffer *bytes.Buffer) {
	buffer.WriteString("network\t")	
}

func Initialize() {
	vms = make(map[string] VirtualMachineExtended)
	models.RegisterCollector(NetworkCollector{})
}
