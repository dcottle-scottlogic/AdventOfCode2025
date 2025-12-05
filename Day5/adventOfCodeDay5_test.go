package adventOfCodeDay5

import "testing"

func TestDay5MainPart1(t *testing.T) {
	got := Day5MainPart1(true, "C:/Dev/AdventOfCode/AdventOfCode2025/Day5/testingInput.txt", "C:/Dev/AdventOfCode/AdventOfCode2025/Day5/testingInput2.txt")
	want := 3
	if got != want {
		t.Errorf("Day2Main() = %v, want %v", got, want)
	}
}

func TestDay5MainPart2(t *testing.T) {
	got := Day5MainPart2("C:/Dev/AdventOfCode/AdventOfCode2025/Day5/testingInput2.txt")
	want := 14
	if got != want {
		t.Errorf("Day2Main() = %v, want %v", got, want)
	}
}
