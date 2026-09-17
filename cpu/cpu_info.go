package cpu

import (
	"os"
	"strconv"
	"strings"

	fileio "github.com/Slynxh07/htop-clone-go/file_io"
)

type CPUInfo struct {
	file       *os.File
	CoreAmount int
}

func NewCPUInfo() (*CPUInfo, error) {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return nil, err
	}

	coreAmountString, err := fileio.ReadLine(13, file)

	coreAmountString = strings.TrimPrefix(coreAmountString, "cpu cores")
	coreAmountString = strings.ReplaceAll(coreAmountString, "\t", "")
	coreAmountString = strings.ReplaceAll(coreAmountString, " ", "")
	coreAmountString = strings.ReplaceAll(coreAmountString, ":", "")

	coreAmount, err := strconv.Atoi(coreAmountString)
	if err != nil {
		return nil, err
	}

	return &CPUInfo{
		file:       file,
		CoreAmount: coreAmount,
	}, nil
}

func (cpu *CPUInfo) Close() error {
	err := cpu.file.Close()
	return err
}
