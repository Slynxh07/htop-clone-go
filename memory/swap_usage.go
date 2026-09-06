package memory

import (
	"log"
	"strconv"
	"strings"

	fileio "github.com/Slynxh07/htop-clone-go/file_io"
)

func GetTotalSwapMem(m MemoryInfo) SwapMemGib {
	totalSwapMemString, err := fileio.ReadLine(15, m.file)
	if err != nil {
		log.Fatal(err)
	}

	totalSwapMemString = strings.TrimPrefix(totalSwapMemString, "SwapTotal: ")
	totalSwapMemString = strings.TrimSuffix(totalSwapMemString, " kB")
	totalSwapMemString = strings.ReplaceAll(totalSwapMemString, " ", "")

	totalSwapMem, err := strconv.Atoi(totalSwapMemString)
	if err != nil {
		log.Fatal(err)
	}

	return SwapMemGib(totalSwapMem) / (1024 * 1024)
}

func CalculateSwapMemUsage(m MemoryInfo, totalSwapMemGib SwapMemGib) SwapMemGib {
	freeSwapMemString, err := fileio.ReadLine(16, m.file)
	if err != nil {
		log.Fatal(err)
	}

	freeSwapMemString = strings.TrimPrefix(freeSwapMemString, "SwapFree: ")
	freeSwapMemString = strings.TrimSuffix(freeSwapMemString, " kB")
	freeSwapMemString = strings.ReplaceAll(freeSwapMemString, " ", "")

	freeSwapMem, err := strconv.Atoi(freeSwapMemString)
	if err != nil {
		log.Fatal(err)
	}

	freeSwapMemGib := SwapMemGib(freeSwapMem) / (1024 * 1024)

	return totalSwapMemGib - freeSwapMemGib
}
