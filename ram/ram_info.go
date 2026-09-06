package ram

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	fileio "github.com/Slynxh07/htop-clone-go/file_io"
)

type RAMInfo struct {
	file        *os.File
	totalMemGib float64
}

func NewRAMInfo() (*RAMInfo, error) {
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

	totalMemGib := float64(totalMem) / (1024 * 1024)

	return &RAMInfo{
		file:        file,
		totalMemGib: totalMemGib,
	}, nil
}

func (r *RAMInfo) Close() error {
	err := r.file.Close()
	if err == nil {
		fmt.Println("Closed /proc/meminfo")
	}
	return err
}
