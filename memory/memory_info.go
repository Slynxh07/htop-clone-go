package memory

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	fileio "github.com/Slynxh07/htop-clone-go/file_io"
)

type SwapMemGib float64
type RAMMemGib float64

type MemoryInfo struct {
	file        *os.File
	totalMemGib RAMMemGib
}

func NewMemoryInfo() (*MemoryInfo, error) {
	file, err := fileio.OpenFile("/proc/meminfo")
	if err != nil {
		return nil, err
	}

	totalMemString, err := fileio.ReadLine(1, file)
	if err != nil {
		log.Fatal(err)
	}

	totalMemString = strings.TrimPrefix(totalMemString, "MemTotal: ")
	totalMemString = strings.TrimSuffix(totalMemString, " kB")
	totalMemString = strings.ReplaceAll(totalMemString, " ", "")

	totalMem, err := strconv.Atoi(totalMemString)
	if err != nil {
		log.Fatal(err)
	}

	totalMemGib := RAMMemGib(totalMem) / (1024 * 1024)

	return &MemoryInfo{
		file:        file,
		totalMemGib: totalMemGib,
	}, nil
}

func (m *MemoryInfo) Close() error {
	err := m.file.Close()
	if err == nil {
		fmt.Println("Closed /proc/meminfo")
	}
	return err
}

func (m MemoryInfo) GetTotalMem() RAMMemGib {
	return m.totalMemGib
}
