package memory

import (
	"github.com/cha87de/kvmtop/models"
	"os/exec"
	"bytes"
	"strings"
	"fmt"
	"strconv"
)

type MemoryCollector struct{
	
}

func (collector MemoryCollector) Collect(vm models.VirtualMachine) (string, error) {
	stats,err := readMemoryStats(vm)
	if err != nil {
		return "-", nil	
	}
	var available, unused int
	if _, exists := stats["unused"]; exists {
		// if unused exists, use values from balloon driver
		available = stats["available"] 
		unused = stats["unused"]
	}else{
		// if it is missing, use actual value, which is inaccurate/wrong since balloon driver is missing
		available = stats["actual"]
		unused = 0		
	}
	
	// byte => mega byte
	var availableMB float64 = float64(available)/1024
	var unusedMB float64 = float64(unused)/1024	
	var usedMB float64 = availableMB-unusedMB
	
	result := fmt.Sprintf("%.0fMB\t%0.fMB", usedMB, availableMB)
	return result, nil
}

func (collector MemoryCollector) CollectDetails(vm models.VirtualMachine) {
	// nothing to do here
}

func DefineFlags() {
	//flag.BoolVar(&CPU_EACH, "cpu-each", CPU_EACH, "CPU each")
}

func PrintHeader(buffer *bytes.Buffer) {
	buffer.WriteString("ram-used\t")
	buffer.WriteString("ram-total\t")	
}

func Initialize() {
	models.RegisterCollector(MemoryCollector{})
}

func readMemoryStats(vm models.VirtualMachine) (map[string]int, error) {
	cmd := exec.Command("virsh", "--connect=qemu:///system", "dommemstat", vm.Name())
	var dommemstat bytes.Buffer
	cmd.Stdout = &dommemstat
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	dommemstatArr := strings.Split(dommemstat.String(), "\n")
	dommemstatFields := make(map[string]int)
	for _, line := range dommemstatArr {
		if(line == ""){
			continue
		}
		lineArr := strings.Split(line, " ")
		val64, err := strconv.ParseInt(lineArr[1], 10, 32)
		if err != nil {
			continue;
		}
		val := int(val64)
		dommemstatFields[lineArr[0]] = val
	}
	
	return dommemstatFields, nil
}