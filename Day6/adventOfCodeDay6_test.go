package adventOfCodeDay6

import "testing"

func TestDay6MainPart1(t *testing.T) {
	got := Day6Main(true, "C:/Dev/AdventOfCode/AdventOfCode2025/Day6/testInput.txt")
	want := 4277556
	if got != want {
		t.Errorf("Day2Main() = %v, want %v", got, want)
	}
}

func TestDay6MainPart2(t *testing.T) {
	got := Day6Main(false, "C:/Dev/AdventOfCode/AdventOfCode2025/Day6/testInput2.txt")
	want := 3263827
	if got != want {
		t.Errorf("Day2Main() = %v, want %v", got, want)
	}
}
