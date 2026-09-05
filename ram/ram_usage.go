package ram

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadMemInfoByLine() {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Println(scanner.Text())
	}
}

func ReadMemInfo() {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
