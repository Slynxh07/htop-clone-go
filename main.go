package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Slynxh07/htop-clone-go/ram"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	ramInfo, err := ram.NewRAMInfo()
	if err != nil {
		log.Fatal(err)
	}
	defer ramInfo.Close()

	for {
		select {
		case sig := <-signals:
			fmt.Println("\nReceived signal:", sig)
			return
		default:
			usedRam := ram.CalculateRAMUsage(*ramInfo)
			fmt.Printf("\r\033[2KRAM In Use: %.2fGiB", usedRam)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
