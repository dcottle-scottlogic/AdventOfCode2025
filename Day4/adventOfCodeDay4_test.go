package adventOfCodeDay4

import "testing"

func TestDay4MainPart1(t *testing.T) {
	got := Day4Main(true, "C:/Dev/AdventOfCode/AdventOfCode2025/Day4/testingInput.txt", 10)
	want := 13
	if got != want {
		t.Errorf("Day2Main() = %v, want %v", got, want)
	}
}

func TestDay4MainPart2(t *testing.T) {
	got := Day4Main(false, "C:/Dev/AdventOfCode/AdventOfCode2025/Day4/testingInput.txt", 10)
	want := 43
	if got != want {
		t.Errorf("Day2Main() = %v, want %v", got, want)
	}
}
