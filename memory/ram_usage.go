package memory

import (
	"log"
	"strconv"
	"strings"

	fileio "github.com/Slynxh07/htop-clone-go/file_io"
)

func CalculateRAMUsage(m MemoryInfo) RAMMemGib {
	availableMemString, err := fileio.ReadLine(3, m.file)
	if err != nil {
		log.Fatal(err)
	}

	availableMemString = strings.TrimPrefix(availableMemString, "MemAvailable: ")
	availableMemString = strings.TrimSuffix(availableMemString, " kB")
	availableMemString = strings.ReplaceAll(availableMemString, " ", "")

	availableMem, err := strconv.Atoi(availableMemString)
	if err != nil {
		log.Fatal(err)
	}

	availableMemGib := RAMMemGib(availableMem) / (1024 * 1024)

	return m.totalMemGib - availableMemGib
}
