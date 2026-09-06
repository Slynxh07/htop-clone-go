package ram

import (
	"log"
	"strconv"
	"strings"

	fileio "github.com/Slynxh07/htop-clone-go/file_io"
)

func CalculateRAMUsage(r RAMInfo) float64 {
	availableMemString, err := fileio.ReadLine(3, r.file)
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

	availableMemGib := float64(availableMem) / (1024 * 1024)

	usedMemGib := r.totalMemGib - availableMemGib

	return usedMemGib
}
