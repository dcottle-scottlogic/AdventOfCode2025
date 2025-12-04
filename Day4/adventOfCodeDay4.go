package adventOfCodeDay4

import (
	"AdventOfCode2025/utils"
)

func GetDay4Input() string {
	return "Day4/dataInput.txt"
}

func Day4Main(isPart1 bool, dataPath string, height int) int {
	ourMap := utils.Generate2DArray(dataPath, height)
	if isPart1 {
		result, _ := checkRows(isPart1, ourMap)
		return result
	}
	result, updatedArray := checkRows(isPart1, ourMap)
	resultsUnchanged := false
	totalAccessibleValues := result
	for !resultsUnchanged {
		result, updatedArray = checkRows(isPart1, updatedArray)
		totalAccessibleValues += result
		if result == 0 {
			resultsUnchanged = true
		}
	}
	return totalAccessibleValues
}

func checkRows(isPart1 bool, ourMap [][]string) (int, [][]string) {
	totalAccessibleRolls := 0
	for rowIndex := 0; rowIndex < len(ourMap); rowIndex++ {
		for columnIndex := 0; columnIndex < len(ourMap[rowIndex]); columnIndex++ {
			currentRowValue := ourMap[rowIndex][columnIndex]
			if currentRowValue == "@" {
				adjacentPaperRolls := 0
				for x := -1; x < 2; x++ {
					for y := -1; y < 2; y++ {
						if x == 0 && y == 0 {
							continue
						}
						adjacentPaperRolls += checkLocation(ourMap, rowIndex+x, columnIndex+y)

					}
				}
				if adjacentPaperRolls < 4 {
					if !isPart1 {
						ourMap[rowIndex][columnIndex] = "X"
					}
					totalAccessibleRolls++
				}
			}
		}
	}
	return totalAccessibleRolls, ourMap
}

func checkLocation(ourMap [][]string, rowIndex int, columnIndex int) int {
	if rowIndex < 0 || columnIndex < 0 || rowIndex >= len(ourMap) || columnIndex >= len(ourMap[0]) {
		return 0
	}
	if ourMap[rowIndex][columnIndex] == "@" {
		return 1
	}
	return 0
}
