package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func FileScanner(filePath string) *bufio.Scanner {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	return bufio.NewScanner(file)
}

func IsEven(num int) bool {
	return num%2 == 0
}

func TimeRun(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Printf("Function took %s", elapsed)
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Could not convert string to int: %v", err)
	}
	return i
}
