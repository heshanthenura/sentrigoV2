package utils

import (
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

type Metrics struct {
	CPUPercent  float64   `json:"cpu_percent"`
	RAMPercent  float64   `json:"ram_percent"`
	CorePercent []float64 `json:"core_percent"`
	DiskIO      uint64    `json:"disk_io"`
}

func GetMetrics() (*Metrics, error) {
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}

	corePercent, err := cpu.Percent(0, true)
	if err != nil {
		return nil, err
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	diskStat, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}

	metrics := &Metrics{
		CPUPercent:  cpuPercent[0],
		RAMPercent:  memInfo.UsedPercent,
		CorePercent: corePercent,
		DiskIO:      diskStat.Used,
	}

	return metrics, nil
}
