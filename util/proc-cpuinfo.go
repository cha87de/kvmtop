package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"kvmtop/config"
)

// ProcCpuinfo represents one entry of cpuinfo proc file
type ProcCpuinfo struct {
	Processor       int
	VendorID        string
	CPUFamily       int
	Model           int
	ModelName       string
	Stepping        int
	Microcode       string
	CPUMhz          float32
	CacheSize       string
	PhysicalID      int
	Siblings        int
	CoreID          int
	CPUCores        int
	ApicID          int
	InitialApicID   int
	Fpu             string
	FpuException    string
	CpuidLevel      int
	Wp              string
	Flags           string
	Bugs            string
	Bogomips        float32
	ClflushSize     int
	CacheAlignment  int
	AddressSizes    string
	PowerManagement string
}

// GetProcCpuinfo reads the proc cpuinfo file
func GetProcCpuinfo() []ProcCpuinfo {
	stats := []ProcCpuinfo{}

	filepath := fmt.Sprint(config.Options.ProcFS, "/cpuinfo")
	/*filecontent, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot read proc cpuinfo: %s\n", err)
		return ProcCpuinfo{}
	}*/

	file, _ := os.Open(filepath)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	core := 0
	for scanner.Scan() {
		row := scanner.Text()
		rowfields := strings.Split(row, ":")

		if len(stats) <= core {
			// create stats for new core
			stats = append(stats, ProcCpuinfo{})
		}
		if len(rowfields) < 2 {
			// new core detected
			core++
			continue
		}

		key := strings.TrimSpace(rowfields[0])
		value := strings.TrimSpace(rowfields[1])
		// fmt.Printf("key: >%s<, value: >%s<\n", key, value)

		switch key {
		case "processor":
			stats[core].Processor, _ = strconv.Atoi(value)
		case "vendor_id":
			stats[core].VendorID = value
		case "cpu family":
			stats[core].CPUFamily, _ = strconv.Atoi(value)
		case "model":
			stats[core].Model, _ = strconv.Atoi(value)
		case "model name":
			stats[core].ModelName = value
		case "stepping":
			stats[core].Stepping, _ = strconv.Atoi(value)
		case "microcode":
			stats[core].Microcode = value
		case "cpu MHz":
			tmp, _ := strconv.ParseFloat(value, 32)
			stats[core].CPUMhz = float32(tmp)
		case "cache size":
			stats[core].CacheSize = value
		case "physical id":
			stats[core].PhysicalID, _ = strconv.Atoi(value)
		case "siblings":
			stats[core].Siblings, _ = strconv.Atoi(value)
		case "core id":
			stats[core].CoreID, _ = strconv.Atoi(value)
		case "cpu cores":
			stats[core].CPUCores, _ = strconv.Atoi(value)
		case "apicid":
			stats[core].ApicID, _ = strconv.Atoi(value)
		case "initial apicid":
			stats[core].InitialApicID, _ = strconv.Atoi(value)
		case "fpu":
			stats[core].Fpu = value
		case "fpu_exception":
			stats[core].FpuException = value
		case "cpuid level":
			stats[core].CpuidLevel, _ = strconv.Atoi(value)
		case "wp":
			stats[core].Wp = value
		case "flags":
			stats[core].Flags = value
		case "bugs":
			stats[core].Bugs = value
		case "bogomips":
			tmp, _ := strconv.ParseFloat(value, 32)
			stats[core].Bogomips = float32(tmp)
		case "clflush size":
			stats[core].ClflushSize, _ = strconv.Atoi(value)
		case "cache_alignment":
			stats[core].CacheAlignment, _ = strconv.Atoi(value)
		case "address sizes":
			stats[core].AddressSizes = value
		case "power management":
			stats[core].PowerManagement = value
		}
	}

	return stats
}
