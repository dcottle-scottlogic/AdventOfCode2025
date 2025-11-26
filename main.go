package main

import (
	"fmt"

	"example.com/AdventOfCode2025/mascot"

	adventOfCodeDay1 "example.com/AdventOfCode2025/Day1"
)

func main() {
	fmt.Println("Welcome to Advent of Code 2025!")
	fmt.Printf("Our mascot is the %s.\n", mascot.GetMascot())
	fmt.Println("----------------------")
	fmt.Println(adventOfCodeDay1.Day1Main())
}
