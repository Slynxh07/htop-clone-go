package memory

import (
	"strconv"
	"strings"

	fileio "github.com/Slynxh07/htop-clone-go/file_io"
)

func GetTotalSwapMem(m MemoryInfo) (SwapMemGib, error) {
	totalSwapMemString, err := fileio.ReadLine(15, m.file)
	if err != nil {
		return 0, err
	}

	totalSwapMemString = strings.TrimPrefix(totalSwapMemString, "SwapTotal: ")
	totalSwapMemString = strings.TrimSuffix(totalSwapMemString, " kB")
	totalSwapMemString = strings.ReplaceAll(totalSwapMemString, " ", "")

	totalSwapMem, err := strconv.Atoi(totalSwapMemString)
	if err != nil {
		return 0, err
	}

	return SwapMemGib(totalSwapMem) / (1024 * 1024), nil
}

func CalculateSwapMemUsage(m MemoryInfo, totalSwapMemGib SwapMemGib) (SwapMemGib, error) {
	freeSwapMemString, err := fileio.ReadLine(16, m.file)
	if err != nil {
		return 0, err
	}

	freeSwapMemString = strings.TrimPrefix(freeSwapMemString, "SwapFree: ")
	freeSwapMemString = strings.TrimSuffix(freeSwapMemString, " kB")
	freeSwapMemString = strings.ReplaceAll(freeSwapMemString, " ", "")

	freeSwapMem, err := strconv.Atoi(freeSwapMemString)
	if err != nil {
		return 0, err
	}

	freeSwapMemGib := SwapMemGib(freeSwapMem) / (1024 * 1024)

	return totalSwapMemGib - freeSwapMemGib, nil
}
