package adventOfCodeDay6

import (
	"AdventOfCode2025/utils"
	"strings"
)

func GetDay6Input() string {
	return "Day6/day6input.txt"
}
func GetDay6Input2() string {
	return "Day6/day6Part2input.txt"
}
func Day6Main(isPart1 bool, path string) int {
	scanner := utils.FileScanner(path)
	array2D := make([][]string, 5)
	row := 0
	runningTotal := 0
	if isPart1 {
		for scanner.Scan() {
			array2D[row] = strings.Split(scanner.Text(), ",")
			row++
		}
		for i := 0; i < len(array2D[0]); i++ {
			int1 := utils.ToInt(array2D[0][i])
			int2 := utils.ToInt(array2D[1][i])
			int3 := utils.ToInt(array2D[2][i])
			int4 := utils.ToInt(array2D[3][i])
			symbol := array2D[4][i]
			if symbol == "*" {
				runningTotal += (int1 * int2 * int3 * int4)
			} else {
				runningTotal += (int1 + int2 + int3 + int4)
			}
		}
	} else {
		for scanner.Scan() {
			array2D[row] = strings.Split(scanner.Text(), "")
			row++
		}
		tempValue1 := ""
		tempValue2 := ""
		tempValue3 := ""
		tempValue4 := ""
		isMult := false
		array2D2 := make([][]string, 5)
		for i := 0; i < len(array2D[0]); i++ {
			if array2D[4][i] == "*" {
				isMult = true
			}
			if array2D[0][i] == " " && array2D[1][i] == " " && array2D[3][i] == " " && array2D[4][i] == " " {
				array2D2[0] = append(array2D2[0], tempValue1)
				array2D2[1] = append(array2D2[1], tempValue2)
				array2D2[2] = append(array2D2[2], tempValue3)
				array2D2[3] = append(array2D2[3], tempValue4)
				if isMult {
					array2D2[4] = append(array2D2[4], "*")
				} else {
					array2D2[4] = append(array2D2[4], "+")
				}
				tempValue1 = ""
				tempValue2 = ""
				tempValue3 = ""
				tempValue4 = ""
				isMult = false
			} else {
				tempValue1 += string(array2D[0][i])
				tempValue2 += string(array2D[1][i])
				tempValue3 += string(array2D[2][i])
				tempValue4 += string(array2D[3][i])
			}
		}
		for i := 0; i < len(array2D2[0]); i++ {
			value1 := array2D2[0][i]
			maxLeng := len(value1)

			value2 := array2D2[1][i]
			value3 := array2D2[2][i]
			value4 := array2D2[3][i]
			symbol := array2D2[4][i]
			ray := make([]string, 0)
			for i := maxLeng - 1; i >= 0; i-- {
				currentValue := ""
				if string(value1[i]) != " " {
					currentValue += string(value1[i])
				}
				if string(value2[i]) != " " {
					currentValue += string(value2[i])

				}
				if string(value3[i]) != " " {
					currentValue += string(value3[i])

				}
				if string(value4[i]) != " " {
					currentValue += string(value4[i])

				}
				ray = append(ray, currentValue)
			}
			rayTotal := utils.ToInt(ray[0])

			if symbol == "*" {
				for i := 1; i < len(ray); i++ {
					rayTotal = rayTotal * utils.ToInt(ray[i])
				}
				runningTotal += rayTotal
			} else {
				for i := 1; i < len(ray); i++ {
					rayTotal = rayTotal + utils.ToInt(ray[i])
				}
				runningTotal += rayTotal
			}
		}
	}
	return runningTotal
}
