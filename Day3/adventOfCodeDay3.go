package adventOfCodeDay3

import (
	"AdventOfCode2025/utils"
	"strings"
)

func GetDay3Input() string {
	return "Day3/day3Input.txt"
}

func Day3Main(isPart1 bool, dataPath string) int {
	if isPart1 {
		return day3Part1(dataPath)
	} else {
		return day3Part2(dataPath)
	}
}

func day3Part1(input string) int {
	totalJoltage := 0
	scanner := utils.FileScanner(input)
	for scanner.Scan() {
		var joltage int = 0
		line := scanner.Text()
		for i := 0; i <= len(line)-1; i++ {
			var tempJoltage int
			for j := i + 1; j < len(line); j++ {
				tempJoltage = utils.ToInt(string(line[i]) + string(line[j]))
				if tempJoltage > joltage {
					joltage = tempJoltage
				}
			}
		}
		totalJoltage += joltage
	}
	return totalJoltage
}

func day3Part2(input string) int {
	totalJoltage := 0
	scanner := utils.FileScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		lineCharArray := strings.Split(line, "")

		for i := 0; len(lineCharArray) > 12; i++ {
			var currentPointer = len(lineCharArray) - 1
			for j := 0; j <= len(lineCharArray)-2; j++ {
				currentInt := utils.ToInt(lineCharArray[j])
				nextInt := utils.ToInt(lineCharArray[j+1])
				if currentInt < nextInt {
					currentPointer = j
					break
				}
			}

			lineCharArray = append(lineCharArray[:currentPointer], lineCharArray[currentPointer+1:]...)
		}

		lineJoltageString := ""
		for k := 0; k < len(lineCharArray); k++ {
			lineJoltageString = lineJoltageString + string(lineCharArray[k])
		}
		totalJoltage += utils.ToInt(lineJoltageString)
	}
	return totalJoltage
}
