package main

import (
	"fmt"

	"github.com/Slynxh07/htop-clone-go/ram"
)

func main() {
	ram.ReadMemInfo()
	fmt.Println()
	ram.ReadMemInfoByLine()
}
