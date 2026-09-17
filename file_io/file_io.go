package fileio

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadLine(line int, file *os.File) (string, error) {
	if line < 1 {
		return "", fmt.Errorf("line number must be greater than 0")
	}

	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)

	currentLine := 1

	for scanner.Scan() {
		if currentLine == line {
			return scanner.Text(), nil
		}

		currentLine++
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("line doesn't exist")
}
