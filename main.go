package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Slynxh07/htop-clone-go/cpu"
	"github.com/Slynxh07/htop-clone-go/memory"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	memInfo, err := memory.NewMemoryInfo()
	if err != nil {
		log.Fatal(err)
	}

	cpuInfo, err := cpu.NewCPUInfo()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("CPU Cores: %d\n", cpuInfo.CoreAmount)

	defer memInfo.Close()
	defer cpuInfo.Close()

	for {
		select {
		case sig := <-signals:
			fmt.Println("\nReceived signal:", sig)
			return
		default:
			totalRam := memInfo.GetTotalMem()
			usedRam, err := memory.CalculateRAMUsage(*memInfo)
			if err != nil {
				log.Fatal(err)
			}
			totalSwap, err := memory.GetTotalSwapMem(*memInfo)
			if err != nil {
				log.Fatal(err)
			}
			usedSwap, err := memory.CalculateSwapMemUsage(*memInfo, totalSwap)
			if err != nil {
				log.Fatal(err)
			}
			swapPercent := 0.0
			if totalSwap > 0 {
				swapPercent = float64(usedSwap/totalSwap) * 100
			}
			fmt.Printf("\r\033[2KRAM: %.2fGiB / %.2fGiB %.2f%%\n", usedRam, totalRam, float64(usedRam/totalRam)*100)
			fmt.Printf("\r\033[2KSwap: %.2fGiB / %.2fGiB %.2f%%", usedSwap, totalSwap, swapPercent)
			fmt.Print("\033[1A")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
