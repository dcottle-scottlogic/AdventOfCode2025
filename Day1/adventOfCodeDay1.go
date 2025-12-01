package adventOfCodeDay1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

//how many times does the dial point to 0

func Day1Main() string {
	file, err := os.Open("Day1/part1Input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	location := 50
	totalDefaults := 0
	for scanner.Scan() {
		var direction string
		var clicks int
		line := scanner.Text()
		splints := strings.Split(line, "L")
		if len(splints) > 1 {
			direction = "Left"
		} else {
			direction = "Right"
			splints = strings.Split(splints[0], "R")
		}
		clicks, error := strconv.Atoi(splints[1])
		if error != nil {
			log.Fatalf("Could not convert to integer: %v", error)
		}
		for i := 0; i < clicks; i++ {
			if direction == "Left" {
				location--
			}
			if direction == "Right" {
				location++
			}
			if location <= -1 {
				location = 99
			}
			if location >= 100 {
				location = 0
			}

		}
		if location == 0 {
			totalDefaults++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error occurred during scanning: %v", err)
	}
	return strconv.Itoa(totalDefaults)
}
