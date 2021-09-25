package main

import (
	"fmt"

	sq "serverquery"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func getSystemInfo() *sq.SystemInformation {
	return &sq.SystemInformation{
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

func getCpuInfo() []*sq.CPUInformation {
	var cpus []*sq.CPUInformation
	c, _ := cpu.Info()
	for _, e := range c {
		entry := &sq.CPUInformation{
			CPU:       e.CPU,
			Vendor:    e.VendorID,
			Family:    e.Family,
			ModelName: e.ModelName,
			Mhz:       e.Mhz,
			Cores:     e.Cores,
			CacheSize: e.CacheSize,
		}
		cpus = append(cpus, entry)
	}
	return cpus
}

func getMemoryInfo() *sq.MemoryInformation {
	v, _ := mem.VirtualMemory()
	return &sq.MemoryInformation{
		Total:       v.Total,
		Available:   v.Available,
		Used:        v.Used,
		UsedPercent: v.UsedPercent,
		Free:        v.Free,
	}
}

func getDiskInfo() *sq.DiskInformation {
	d, _ := disk.Usage("/")
	return &sq.DiskInformation{
		Total:       d.Total,
		Free:        d.Free,
		Used:        d.Used,
		UsedPercent: d.UsedPercent,
	}
}
