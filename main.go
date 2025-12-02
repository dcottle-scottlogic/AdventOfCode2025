package main

import (
	"fmt"

	adventOfCodeDay1 "AdventOfCode2025/Day1"
	adventOfCodeDay2 "AdventOfCode2025/Day2"
	"AdventOfCode2025/mascot"
)

func main() {
	fmt.Println("Welcome to Advent of Code 2025!")
	var ourMascot = mascot.GetMascot()
	fmt.Printf("Our mascot is the %s a %s.\n", ourMascot.GetMascotName(), ourMascot.GetMascotType())
	fmt.Println("----------------------")
	fmt.Println("Day 1 Solution:")
	fmt.Println("The result for part one is: " + adventOfCodeDay1.Day1Main(true))
	fmt.Println("The result for part two is: " + adventOfCodeDay1.Day1Main(false))
	fmt.Println("----------------------")
	fmt.Println("Day 2 Solution:")
	fmt.Println("The result for part one is: " + adventOfCodeDay2.Day2Main(true, adventOfCodeDay2.GetDay2Input()))
	fmt.Println("The result for part two is: " + adventOfCodeDay2.Day2Main(false, adventOfCodeDay2.GetDay2Input()))
}
