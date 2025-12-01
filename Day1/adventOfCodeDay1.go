package adventOfCodeDay1

import (
	"log"
	"strconv"
	"strings"

	"example.com/AdventOfCode2025/utils"
)

//how many times does the dial point to 0

func Day1Main() string {
	scanner := utils.FileScanner("Day1/part1Input.txt")
	location := 50
	totalDefaults := 0
	for scanner.Scan() { // Read each line
		var direction string = ""
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
			} else {
				location++
			}
			if location <= -1 {
				location = 99
			}
			if location >= 100 {
				location = 0
			}
			//For part 2 only
			if location == 0 {
				totalDefaults++
			}

		}
		//For part 1 only
		// if location == 0 {
		// 	totalDefaults++
		// }

		//Part 1 answer = 1074
		//Part 2 answer = 6254
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error occurred during scanning: %v", err)
	}
	return strconv.Itoa(totalDefaults)
}
