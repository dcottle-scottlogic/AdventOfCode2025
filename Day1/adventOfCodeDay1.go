package adventOfCodeDay1

import (
	"log"
	"strconv"

	"example.com/AdventOfCode2025/utils"
)

func Day1Main(isPart1 bool) string {

	scanner := utils.FileScanner("Day1/part1Input.txt")
	location := 50
	totalDefaults := 0

	for scanner.Scan() {
		var isCounterClockwise bool
		var clicks int
		line := scanner.Text()
		clickString := line[1:]

		if line[0] == 'L' {
			isCounterClockwise = true
		} else {
			isCounterClockwise = false
		}

		clicks, error := strconv.Atoi(clickString)
		if error != nil {
			log.Fatalf("Could not convert to integer: %v", error)
		}

		for i := 0; i < clicks; i++ {
			if isCounterClockwise {
				location--
			} else {
				location++
			}
			if location <= -1 {
				location = 99
			}
			if location >= 100 {
				location = 0
			}
			if location == 0 && !isPart1 {
				totalDefaults++
			}

		}
		if location == 0 && isPart1 {
			totalDefaults++
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error occurred during scanning: %v", err)
	}
	return strconv.Itoa(totalDefaults)
}
