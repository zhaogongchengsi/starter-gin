package utils

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"runtime"
	"time"
)

const (
	Byte = 1
	KByt = 1024 * Byte
	MByt = 1024 * KByt
	GByt = 1024 * MByt
)

type Os struct {
	GOOS         string `json:"goos"`
	NumCPU       int    `json:"numCpu"`
	Compiler     string `json:"compiler"`
	GoVersion    string `json:"goVersion"`
	NumGoroutine int    `json:"numGoroutine"`
}

type Cpu struct {
	Cpus  []float64 `json:"cpus"`
	Cores int       `json:"cores"`
}

type Ram struct {
	TotalMB     int `json:"totalMb"`
	UsedMB      int `json:"usedMb"`
	UsedPercent int `json:"usedPercent"`
}

type Disk struct {
	UsedMB      int `json:"usedMb"`
	TotalGB     int `json:"totalGb"`
	UsedGB      int `json:"usedGb"`
	TotalMB     int `json:"totalMb"`
	UsedPercent int `json:"usedPercent"`
}

type Health struct {
	Os   `json:"os"`
	Cpu  `json:"cpu"`
	Ram  `json:"ram"`
	Disk `json:"disk"`
}

func NewOs() *Os {
	return &Os{
		GOOS:         runtime.GOOS,
		Compiler:     runtime.Compiler,
		NumCPU:       runtime.NumCPU(),
		GoVersion:    runtime.Version(),
		NumGoroutine: runtime.NumGoroutine(),
	}
}

func NewCpu() (c Cpu, err error) {
	if cores, err := cpu.Counts(false); err != nil {
		return c, err
	} else {
		c.Cores = cores
	}
	if cpus, err := cpu.Percent(time.Duration(200)*time.Millisecond, true); err != nil {
		return c, err
	} else {
		c.Cpus = cpus
	}
	return c, nil
}

func NewRAM() (r Ram, err error) {
	if u, err := mem.VirtualMemory(); err != nil {
		return r, err
	} else {
		return Ram{
			UsedMB:      int(u.Used) / MByt,
			TotalMB:     int(u.Total) / MByt,
			UsedPercent: int(u.UsedPercent),
		}, nil
	}
}

func NewDisk() (d Disk, err error) {
	if u, err := disk.Usage("/"); err != nil {
		return d, err
	} else {
		return Disk{
			UsedMB:      int(u.Used) / MByt,
			UsedGB:      int(u.Used) / GByt,
			TotalMB:     int(u.Total) / MByt,
			TotalGB:     int(u.Total) / MByt,
			UsedPercent: int(u.UsedPercent) / MByt,
		}, nil
	}
}
