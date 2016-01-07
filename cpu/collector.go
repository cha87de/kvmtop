package cpu

import (
	"fmt"
	//"flag"
	"bytes"
	"github.com/cha87de/kvmtop/models"
)

var (
	// store stats for calculating diffs
	cpustats map[string]VmCpuSchedStat
	CPU_EACH = false
)

type CpuCollector struct {
}

func (collector CpuCollector) Collect(vm models.VirtualMachine) (string, error) {
	// new current CPU counters
	newSchedStat, err := lookupStats(vm)
	if err != nil {
		return "", err
	}

	// get old CPU counters
	if oldSchedStat, exists := cpustats[vm.Ppid()]; exists {
		// set new stats as old stats for next run
		cpustats[vm.Ppid()] = newSchedStat
		// calculate diff between new and old counters
		cpu_utilisation := newSchedStat.calculateDiff(oldSchedStat)
		vCores := len(vm.VCpuTasks())
		result := fmt.Sprintf("%d\t%.0f%%\t%.0f%%\t%.0f%%\t%.0f%%", 
					vCores, 
					(cpu_utilisation.Avg.Inside * 100), 
					(cpu_utilisation.Avg.Outside * 100), 
					(cpu_utilisation.Avg.Steal * 100), 
					(cpu_utilisation.Avg_other * 100))
		return result, nil
	} else {
		// no measurement yet
		// set new stats as old stats for next run
		cpustats[vm.Ppid()] = newSchedStat
		return "-", nil
	}
}

func (collector CpuCollector) CollectDetails(vm models.VirtualMachine) {
	// TODO lookup vCores here! not in VirtualMachine
}

func DefineFlags() {
	//flag.BoolVar(&CPU_EACH, "cpu-each", CPU_EACH, "CPU each")
}

func PrintHeader(buffer *bytes.Buffer) {
	buffer.WriteString("CpuCS\t")
	buffer.WriteString("CpuVM\t")
	buffer.WriteString("CpuPM\t")
	buffer.WriteString("CpuST\t")
	buffer.WriteString("CpuIO\t")
}

func Initialize() {
	if CPU_EACH {
		fmt.Println("PRINT EACH CPU")
	}

	cpustats = make(map[string]VmCpuSchedStat)
	models.RegisterCollector(CpuCollector{})
}
