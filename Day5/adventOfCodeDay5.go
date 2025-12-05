package adventOfCodeDay5

import (
	"AdventOfCode2025/utils"
	"strings"
)

func GetDay5Input1() string {
	return "Day5/dataInputIds.txt"
}
func GetDay5Input2() string {
	return "Day5/dataInputRanges.txt"
}

func Day5MainPart1(isPart1 bool, ingredientIds string, ranges string) int {
	count := 0
	ingredientScanner := utils.FileScanner(ingredientIds)
	for ingredientScanner.Scan() {
		id := utils.ToInt(ingredientScanner.Text())
		rangeScanner := utils.FileScanner(ranges)
		for rangeScanner.Scan() {
			rangeLine := strings.Split(rangeScanner.Text(), "-")
			if id >= utils.ToInt(rangeLine[0]) && id <= utils.ToInt(rangeLine[1]) {
				count++
				break
			}
		}
	}
	return count
}

func Day5MainPart2(ranges string) int {
	rangeScanner := utils.FileScanner(ranges)
	fresh := [][2]int{}

	for rangeScanner.Scan() {
		newSplits := strings.Split(rangeScanner.Text(), "-")
		newList := [2]int{utils.ToInt(newSplits[0]), utils.ToInt(newSplits[1])}
		_, fresh = squashRanges(newList, fresh)
	}

	changed := true //condense further
	for changed {
		changed = false
		newList := [][2]int{}

		for _, r := range fresh {
			merged, out := squashRanges(r, newList)
			newList = out
			if merged {
				changed = true
			}
		}

		fresh = newList
	}

	total := 0
	for _, r := range fresh {
		total += (r[1] - r[0]) + 1 // plus 1 to be inclusive
	}

	return total
}

func squashRanges(newRange [2]int, twoDArray [][2]int) (bool, [][2]int) {
	start := newRange[0]
	end := newRange[1]

	newList := [][2]int{}
	merged := false

	for _, existing := range twoDArray {
		existingStartId := existing[0]
		existingEndId := existing[1]

		if !(end < existingStartId-1 || start > existingEndId+1) {
			start = min(start, existingStartId)
			end = max(end, existingEndId)
			merged = true
		} else {
			newList = append(newList, existing)
		}
	}

	newList = append(newList, [2]int{start, end})
	return merged, newList
}
