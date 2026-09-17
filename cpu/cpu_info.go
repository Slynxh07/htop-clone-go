package cpu

import (
	"fmt"
	"os"
)

type CPUInfo struct {
	file        *os.File
	coresAmount int
}

func NewCPUInfo() (*CPUInfo, error) {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return nil, err
	}

	return &CPUInfo{
		file:        file,
		coresAmount: 8,
	}, nil
}

func (cpu *CPUInfo) Close() error {
	err := cpu.file.Close()
	if err == nil {
		fmt.Println("Closed /proc/cpuinfo")
	}
	return err
}
