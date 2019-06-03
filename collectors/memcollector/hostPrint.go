package memcollector

import (
	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
	"fmt"
)

func hostPrint(host *models.Host) []string {

	Total, err := host.GetMetricUint64("ram_Total", 0)
	if err != nil {
		fmt.Printf("error getting ram_total: %s\n", err)
	}
	Free, _ := host.GetMetricUint64("ram_Free", 0)
	Available,_ := host.GetMetricUint64("ram_Available", 0)
	Buffers,_ := host.GetMetricUint64("ram_Buffers", 0)
	Cached,_ := host.GetMetricUint64("ram_Cached", 0)
	SwapCached,_ := host.GetMetricUint64("ram_SwapCached", 0)
	Active,_ := host.GetMetricUint64("ram_Active", 0)
	Inactive,_ := host.GetMetricUint64("ram_Inactive", 0)
	ActiveAanon,_ := host.GetMetricUint64("ram_ActiveAanon", 0)
	InactiveAanon,_ := host.GetMetricUint64("ram_InactiveAanon", 0)
	ActiveFile,_ := host.GetMetricUint64("ram_ActiveFile", 0)
	InactiveFile,_ := host.GetMetricUint64("ram_InactiveFile", 0)
	Unevictable,_ := host.GetMetricUint64("ram_Unevictable", 0)
	Mlocked,_ := host.GetMetricUint64("ram_Mlocked", 0)
	SwapTotal,_ := host.GetMetricUint64("ram_SwapTotal", 0)
	SwapFree,_ := host.GetMetricUint64("ram_SwapFree", 0)
	Dirty,_ := host.GetMetricUint64("ram_Dirty", 0)
	Writeback,_ := host.GetMetricUint64("ram_Writeback", 0)
	AnonPages,_ := host.GetMetricUint64("ram_AnonPages", 0)
	Mapped,_ := host.GetMetricUint64("ram_Mapped", 0)
	Shmem,_ := host.GetMetricUint64("ram_Shmem", 0)
	Slab,_ := host.GetMetricUint64("ram_Slab", 0)
	SReclaimable,_ := host.GetMetricUint64("ram_SReclaimable", 0)
	SUnreclaim,_ := host.GetMetricUint64("ram_SUnreclaim", 0)
	KernelStack,_ := host.GetMetricUint64("ram_KernelStack", 0)
	PageTables,_ := host.GetMetricUint64("ram_PageTables", 0)
	NFSUnstable,_ := host.GetMetricUint64("ram_NFSUnstable", 0)
	Bounce,_ := host.GetMetricUint64("ram_Bounce", 0)
	WritebackTmp,_ := host.GetMetricUint64("ram_WritebackTmp", 0)
	CommitLimit,_ := host.GetMetricUint64("ram_CommitLimit", 0)
	CommittedAS,_ := host.GetMetricUint64("ram_CommittedAS", 0)
	VmallocTotal,_ := host.GetMetricUint64("ram_VmallocTotal", 0)
	VmallocUsed,_ := host.GetMetricUint64("ram_VmallocUsed", 0)
	VmallocChunk,_ := host.GetMetricUint64("ram_VmallocChunk", 0)
	HardwareCorrupted,_ := host.GetMetricUint64("ram_HardwareCorrupted", 0)
	AnonHugePages,_ := host.GetMetricUint64("ram_AnonHugePages", 0)
	ShmemHugePages,_ := host.GetMetricUint64("ram_ShmemHugePages", 0)
	ShmemPmdMapped,_ := host.GetMetricUint64("ram_ShmemPmdMapped", 0)
	HugePagesTotal,_ := host.GetMetricUint64("ram_HugePagesTotal", 0)
	HugePagesFree,_ := host.GetMetricUint64("ram_HugePagesFree", 0)
	HugePagesRsvd,_ := host.GetMetricUint64("ram_HugePagesRsvd", 0)
	HugePagesSurp,_ := host.GetMetricUint64("ram_HugePagesSurp", 0)
	Hugepagesize,_ := host.GetMetricUint64("ram_Hugepagesize", 0)
	Hugetlb,_ := host.GetMetricUint64("ram_Hugetlb", 0)
	DirectMap4k,_ := host.GetMetricUint64("ram_DirectMap4k", 0)
	DirectMap2M,_ := host.GetMetricUint64("ram_DirectMap2M", 0)
	DirectMap1G,_ := host.GetMetricUint64("ram_DirectMap1G", 0)

	result := append([]string{Total}, Free, Available)
	if config.Options.Verbose {
		result = append(result, Buffers, Cached, SwapCached, Active, Inactive, ActiveAanon, InactiveAanon, ActiveFile, InactiveFile, Unevictable, Mlocked, SwapTotal, SwapFree, Dirty, Writeback, AnonPages, Mapped, Shmem, Slab, SReclaimable, SUnreclaim, KernelStack, PageTables, NFSUnstable, Bounce, WritebackTmp, CommitLimit, CommittedAS, VmallocTotal, VmallocUsed, VmallocChunk, HardwareCorrupted, AnonHugePages, ShmemHugePages, ShmemPmdMapped, HugePagesTotal, HugePagesFree, HugePagesRsvd, HugePagesSurp, Hugepagesize, Hugetlb, DirectMap4k, DirectMap2M, DirectMap1G)
	}
	return result

}
