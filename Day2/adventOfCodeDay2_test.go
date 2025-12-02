package adventOfCodeDay2

import "testing"

var exampleList = []string{"38593856-38593862", "446443-446449", "1698522-1698528", "222220-222224", "1188511880-1188511890", "998-1012", "95-115", "11-22", "565653-565659",
	"824824821-824824827", "2121212118-2121212124"}

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
