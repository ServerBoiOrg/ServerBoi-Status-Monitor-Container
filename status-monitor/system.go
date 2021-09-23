package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemInformation struct {
	Uptime string             `json:"Uptime"`
	CPU    []*CPUInformation  `json"CPU"`
	Memory *MemoryInformation `json"Memory"`
	Disk   *DiskInformation   `json:"Disk"`
}

type MemoryInformation struct {
	Total       uint64  `json:"Total"`
	Available   uint64  `json:"Available"`
	Used        uint64  `json:"Used"`
	UsedPercent float64 `json:"Percent-Used"`
	Free        uint64  `json:"Free"`
}

type DiskInformation struct {
	Fstype      string  `json:"File-System"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"Free"`
	Used        uint64  `json:"Used"`
	UsedPercent float64 `json:"Percent-Used"`
}

type CPUInformation struct {
	CPU       int32   `json:"Cpu"`
	Vendor    string  `json:"Vendor"`
	Family    string  `json:"Family"`
	Model     string  `json:"Model"`
	CoreID    string  `json:"Core-ID"`
	Cores     int32   `json:"Cores"`
	ModelName string  `json:"Model-Name"`
	Mhz       float64 `json:"Mhz"`
	CacheSize int32   `json:"Cache=Size"`
}

func getSystemInfo() *SystemInformation {
	return &SystemInformation{
		Uptime: getUptime(),
		CPU:    getCpuInfo(),
		Disk:   getDiskInfo(),
		Memory: getMemoryInfo(),
	}
}

func getUptime() string {
	uptime, _ := host.Uptime()
	days := uptime / (60 * 60 * 24)
	hours := (uptime - (days * 60 * 60 * 24)) / (60 * 60)
	minutes := ((uptime - (days * 60 * 60 * 24)) - (hours * 60 * 60)) / 60
	return fmt.Sprintf("%d days, %d hours, %d minutes", days, hours, minutes)
}

func getCpuInfo() []*CPUInformation {
	var cpus []*CPUInformation
	c, _ := cpu.Info()
	for _, e := range c {
		entry := &CPUInformation{
			CPU:       e.CPU,
			Vendor:    e.VendorID,
			Family:    e.Family,
			Model:     e.Model,
			ModelName: e.ModelName,
			Mhz:       e.Mhz,
			CoreID:    e.CoreID,
			Cores:     e.Cores,
			CacheSize: e.CacheSize,
		}
		cpus = append(cpus, entry)
	}
	return cpus
}

func getMemoryInfo() *MemoryInformation {
	v, _ := mem.VirtualMemory()
	return &MemoryInformation{
		Total:       v.Total,
		Available:   v.Available,
		Used:        v.Used,
		UsedPercent: v.UsedPercent,
	}
}

func getDiskInfo() *DiskInformation {
	d, _ := disk.Usage("/")
	return &DiskInformation{
		Fstype:      d.Fstype,
		Total:       d.Total,
		Free:        d.Free,
		Used:        d.Used,
		UsedPercent: d.UsedPercent,
	}
}
