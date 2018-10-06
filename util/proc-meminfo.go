package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cha87de/kvmtop/config"
)

// ProcMeminfo represents the content of the meminfo proc file
type ProcMeminfo struct {
	MemTotal          int
	MemFree           int
	MemAvailable      int
	Buffers           int
	Cached            int
	SwapCached        int
	Active            int
	Inactive          int
	ActiveAanon       int
	InactiveAanon     int
	ActiveFile        int
	InactiveFile      int
	Unevictable       int
	Mlocked           int
	SwapTotal         int
	SwapFree          int
	Dirty             int
	Writeback         int
	AnonPages         int
	Mapped            int
	Shmem             int
	Slab              int
	SReclaimable      int
	SUnreclaim        int
	KernelStack       int
	PageTables        int
	NFSUnstable       int
	Bounce            int
	WritebackTmp      int
	CommitLimit       int
	CommittedAS       int
	VmallocTotal      int
	VmallocUsed       int
	VmallocChunk      int
	HardwareCorrupted int
	AnonHugePages     int
	ShmemHugePages    int
	ShmemPmdMapped    int
	HugePagesTotal    int
	HugePagesFree     int
	HugePagesRsvd     int
	HugePagesSurp     int
	Hugepagesize      int
	Hugetlb           int
	DirectMap4k       int
	DirectMap2M       int
	DirectMap1G       int
}

// GetProcMeminfo reads the proc cpuinfo file
func GetProcMeminfo() ProcMeminfo {
	stats := ProcMeminfo{}

	filepath := fmt.Sprint(config.Options.ProcFS, "/meminfo")

	file, _ := os.Open(filepath)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		row := scanner.Text()
		rowfields := strings.Split(row, ":")

		key := strings.TrimSpace(rowfields[0])
		valueRaw := strings.TrimSpace(rowfields[1])
		valueParts := strings.Split(valueRaw, " ")
		value := valueParts[0]

		switch key {
		case "MemTotal":
			stats.MemTotal, _ = strconv.Atoi(value)
		case "MemFree":
			stats.MemFree, _ = strconv.Atoi(value)
		case "MemAvailable":
			stats.MemAvailable, _ = strconv.Atoi(value)
		case "Buffers":
			stats.Buffers, _ = strconv.Atoi(value)
		case "Cached":
			stats.Cached, _ = strconv.Atoi(value)
		case "SwapCached":
			stats.SwapCached, _ = strconv.Atoi(value)
		case "Active":
			stats.Active, _ = strconv.Atoi(value)
		case "Inactive":
			stats.Inactive, _ = strconv.Atoi(value)
		case "Active(anon)":
			stats.ActiveAanon, _ = strconv.Atoi(value)
		case "Inactive(anon)":
			stats.InactiveAanon, _ = strconv.Atoi(value)
		case "Active(file)":
			stats.ActiveFile, _ = strconv.Atoi(value)
		case "Inactive(file)":
			stats.InactiveFile, _ = strconv.Atoi(value)
		case "Unevictable":
			stats.Unevictable, _ = strconv.Atoi(value)
		case "Mlocked":
			stats.Mlocked, _ = strconv.Atoi(value)
		case "SwapTotal":
			stats.SwapTotal, _ = strconv.Atoi(value)
		case "SwapFree":
			stats.SwapFree, _ = strconv.Atoi(value)
		case "Dirty":
			stats.Dirty, _ = strconv.Atoi(value)
		case "Writeback":
			stats.Writeback, _ = strconv.Atoi(value)
		case "AnonPages":
			stats.AnonPages, _ = strconv.Atoi(value)
		case "Mapped":
			stats.Mapped, _ = strconv.Atoi(value)
		case "Shmem":
			stats.Shmem, _ = strconv.Atoi(value)
		case "Slab":
			stats.Slab, _ = strconv.Atoi(value)
		case "SReclaimable":
			stats.SReclaimable, _ = strconv.Atoi(value)
		case "SUnreclaim":
			stats.SUnreclaim, _ = strconv.Atoi(value)
		case "KernelStack":
			stats.KernelStack, _ = strconv.Atoi(value)
		case "PageTables":
			stats.PageTables, _ = strconv.Atoi(value)
		case "NFS_Unstable":
			stats.NFSUnstable, _ = strconv.Atoi(value)
		case "Bounce":
			stats.Bounce, _ = strconv.Atoi(value)
		case "WritebackTmp":
			stats.WritebackTmp, _ = strconv.Atoi(value)
		case "CommitLimit":
			stats.CommitLimit, _ = strconv.Atoi(value)
		case "Committed_AS":
			stats.CommittedAS, _ = strconv.Atoi(value)
		case "VmallocTotal":
			stats.VmallocTotal, _ = strconv.Atoi(value)
		case "VmallocUsed":
			stats.VmallocUsed, _ = strconv.Atoi(value)
		case "VmallocChunk":
			stats.VmallocChunk, _ = strconv.Atoi(value)
		case "HardwareCorrupted":
			stats.HardwareCorrupted, _ = strconv.Atoi(value)
		case "AnonHugePages":
			stats.AnonHugePages, _ = strconv.Atoi(value)
		case "ShmemHugePages":
			stats.ShmemHugePages, _ = strconv.Atoi(value)
		case "ShmemPmdMapped":
			stats.ShmemPmdMapped, _ = strconv.Atoi(value)
		case "HugePages_Total":
			stats.HugePagesTotal, _ = strconv.Atoi(value)
		case "HugePages_Free":
			stats.HugePagesFree, _ = strconv.Atoi(value)
		case "HugePages_Rsvd":
			stats.HugePagesRsvd, _ = strconv.Atoi(value)
		case "HugePages_Surp":
			stats.HugePagesSurp, _ = strconv.Atoi(value)
		case "Hugepagesize":
			stats.Hugepagesize, _ = strconv.Atoi(value)
		case "Hugetlb":
			stats.Hugetlb, _ = strconv.Atoi(value)
		case "DirectMap4k":
			stats.DirectMap4k, _ = strconv.Atoi(value)
		case "DirectMap2M":
			stats.DirectMap2M, _ = strconv.Atoi(value)
		case "DirectMap1G":
			stats.DirectMap1G, _ = strconv.Atoi(value)
		}
	}

	return stats
}
