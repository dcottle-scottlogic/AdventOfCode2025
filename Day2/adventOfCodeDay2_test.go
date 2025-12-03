package adventOfCodeDay2

import "testing"

var exampleList = []string{}

func TestDay2MainPart1(t *testing.T) {
	got := Day2Main(true, exampleList)
	want := "1227775554"
	if got != want {
		t.Errorf("Day2Main() = %q, want %q", got, want)
	}
}

func TestDay2MainPart2(t *testing.T) {
	got := Day2Main(false, exampleList)
	want := "4174379265"
	if got != want {
		t.Errorf("Day2Main() = %q, want %q", got, want)
	}
}
