package adventOfCodeDay3

import "testing"

func TestDay3MainPart1(t *testing.T) {
	got := Day3Main(true, "C:/Dev/AdventOfCode/AdventOfCode2025/Day3/testData.txt")
	want := 357
	if got != want {
		t.Errorf("Day2Main() = %v, want %v", got, want)
	}
}

func TestDay3MainPart2(t *testing.T) {
	got := Day3Main(false, "C:/Dev/AdventOfCode/AdventOfCode2025/Day3/testData.txt")
	want := 3121910778619
	if got != want {
		t.Errorf("Day2Main() = %v, want %v", got, want)
	}
}
