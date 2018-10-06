package memcollector

import (
	"github.com/cha87de/kvmtop/collectors"
	"github.com/cha87de/kvmtop/models"
)

func hostPrint(host *models.Host) []string {

	Total := collectors.GetMetricUint64(host.Measurable, "ram_Total", 0)
	Free := collectors.GetMetricUint64(host.Measurable, "ram_Free", 0)
	Available := collectors.GetMetricUint64(host.Measurable, "ram_Available", 0)
	Buffers := collectors.GetMetricUint64(host.Measurable, "ram_Buffers", 0)
	Cached := collectors.GetMetricUint64(host.Measurable, "ram_Cached", 0)
	SwapCached := collectors.GetMetricUint64(host.Measurable, "ram_SwapCached", 0)
	Active := collectors.GetMetricUint64(host.Measurable, "ram_Active", 0)
	Inactive := collectors.GetMetricUint64(host.Measurable, "ram_Inactive", 0)
	ActiveAanon := collectors.GetMetricUint64(host.Measurable, "ram_ActiveAanon", 0)
	InactiveAanon := collectors.GetMetricUint64(host.Measurable, "ram_InactiveAanon", 0)
	ActiveFile := collectors.GetMetricUint64(host.Measurable, "ram_ActiveFile", 0)
	InactiveFile := collectors.GetMetricUint64(host.Measurable, "ram_InactiveFile", 0)
	Unevictable := collectors.GetMetricUint64(host.Measurable, "ram_Unevictable", 0)
	Mlocked := collectors.GetMetricUint64(host.Measurable, "ram_Mlocked", 0)
	SwapTotal := collectors.GetMetricUint64(host.Measurable, "ram_SwapTotal", 0)
	SwapFree := collectors.GetMetricUint64(host.Measurable, "ram_SwapFree", 0)
	Dirty := collectors.GetMetricUint64(host.Measurable, "ram_Dirty", 0)
	Writeback := collectors.GetMetricUint64(host.Measurable, "ram_Writeback", 0)
	AnonPages := collectors.GetMetricUint64(host.Measurable, "ram_AnonPages", 0)
	Mapped := collectors.GetMetricUint64(host.Measurable, "ram_Mapped", 0)
	Shmem := collectors.GetMetricUint64(host.Measurable, "ram_Shmem", 0)
	Slab := collectors.GetMetricUint64(host.Measurable, "ram_Slab", 0)
	SReclaimable := collectors.GetMetricUint64(host.Measurable, "ram_SReclaimable", 0)
	SUnreclaim := collectors.GetMetricUint64(host.Measurable, "ram_SUnreclaim", 0)
	KernelStack := collectors.GetMetricUint64(host.Measurable, "ram_KernelStack", 0)
	PageTables := collectors.GetMetricUint64(host.Measurable, "ram_PageTables", 0)
	NFSUnstable := collectors.GetMetricUint64(host.Measurable, "ram_NFSUnstable", 0)
	Bounce := collectors.GetMetricUint64(host.Measurable, "ram_Bounce", 0)
	WritebackTmp := collectors.GetMetricUint64(host.Measurable, "ram_WritebackTmp", 0)
	CommitLimit := collectors.GetMetricUint64(host.Measurable, "ram_CommitLimit", 0)
	CommittedAS := collectors.GetMetricUint64(host.Measurable, "ram_CommittedAS", 0)
	VmallocTotal := collectors.GetMetricUint64(host.Measurable, "ram_VmallocTotal", 0)
	VmallocUsed := collectors.GetMetricUint64(host.Measurable, "ram_VmallocUsed", 0)
	VmallocChunk := collectors.GetMetricUint64(host.Measurable, "ram_VmallocChunk", 0)
	HardwareCorrupted := collectors.GetMetricUint64(host.Measurable, "ram_HardwareCorrupted", 0)
	AnonHugePages := collectors.GetMetricUint64(host.Measurable, "ram_AnonHugePages", 0)
	ShmemHugePages := collectors.GetMetricUint64(host.Measurable, "ram_ShmemHugePages", 0)
	ShmemPmdMapped := collectors.GetMetricUint64(host.Measurable, "ram_ShmemPmdMapped", 0)
	HugePagesTotal := collectors.GetMetricUint64(host.Measurable, "ram_HugePagesTotal", 0)
	HugePagesFree := collectors.GetMetricUint64(host.Measurable, "ram_HugePagesFree", 0)
	HugePagesRsvd := collectors.GetMetricUint64(host.Measurable, "ram_HugePagesRsvd", 0)
	HugePagesSurp := collectors.GetMetricUint64(host.Measurable, "ram_HugePagesSurp", 0)
	Hugepagesize := collectors.GetMetricUint64(host.Measurable, "ram_Hugepagesize", 0)
	Hugetlb := collectors.GetMetricUint64(host.Measurable, "ram_Hugetlb", 0)
	DirectMap4k := collectors.GetMetricUint64(host.Measurable, "ram_DirectMap4k", 0)
	DirectMap2M := collectors.GetMetricUint64(host.Measurable, "ram_DirectMap2M", 0)
	DirectMap1G := collectors.GetMetricUint64(host.Measurable, "ram_DirectMap1G", 0)

	result := append([]string{Total}, Free, Available, Buffers, Cached, SwapCached, Active, Inactive, ActiveAanon, InactiveAanon, ActiveFile, InactiveFile, Unevictable, Mlocked, SwapTotal, SwapFree, Dirty, Writeback, AnonPages, Mapped, Shmem, Slab, SReclaimable, SUnreclaim, KernelStack, PageTables, NFSUnstable, Bounce, WritebackTmp, CommitLimit, CommittedAS, VmallocTotal, VmallocUsed, VmallocChunk, HardwareCorrupted, AnonHugePages, ShmemHugePages, ShmemPmdMapped, HugePagesTotal, HugePagesFree, HugePagesRsvd, HugePagesSurp, Hugepagesize, Hugetlb, DirectMap4k, DirectMap2M, DirectMap1G)
	return result

}
