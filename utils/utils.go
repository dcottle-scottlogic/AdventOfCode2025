package utils

import (
	"bufio"
	"log"
	"os"
)

func FileScanner(filePath string) *bufio.Scanner {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	return bufio.NewScanner(file)
}
