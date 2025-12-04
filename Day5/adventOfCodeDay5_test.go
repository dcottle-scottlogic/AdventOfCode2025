package adventOfCodeDay5

import "testing"

func TestDay5MainPart1(t *testing.T) {
	got := Day5Main(true, "C:/Dev/AdventOfCode/AdventOfCode2025/Day5/testingInput.txt")
	want := 0
	if got != want {
		t.Errorf("Day2Main() = %v, want %v", got, want)
	}
}

func TestDay5MainPart2(t *testing.T) {
	got := Day5Main(false, "C:/Dev/AdventOfCode/AdventOfCode2025/Day5/testingInput.txt")
	want := 0
	if got != want {
		t.Errorf("Day2Main() = %v, want %v", got, want)
	}
}
