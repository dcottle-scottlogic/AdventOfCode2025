package adventOfCodeDay2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"example.com/AdventOfCode2025/utils"
)

var exampleList = []string{"38593856-38593862", "446443-446449", "1698522-1698528", "222220-222224", "1188511880-1188511890", "998-1012", "95-115", "11-22", "565653-565659",
	"824824821-824824827", "2121212118-2121212124"}
var inputList = []string{"18623-26004", "226779-293422", "65855-88510", "868-1423", "248115026-248337139", "903911-926580", "97-121", "67636417-67796062", "24-47", "6968-10197", "193-242", "3769-5052", "5140337-5233474", "2894097247-2894150301", "979582-1016336", "502-646", "9132195-9191022", "266-378", "58-91", "736828-868857", "622792-694076", "6767592127-6767717303", "2920-3656", "8811329-8931031", "107384-147042", "941220-969217", "3-17", "360063-562672", "7979763615-7979843972", "1890-2660", "23170346-23308802"}

func Day2Main(isPart1 bool) string {
	var InvalidRunningTotal int = 0
	fmt.Println("Starting Day 2 Computation")
	for i := range inputList {
		var idPair string = inputList[i]
		if isPart1 {
			InvalidRunningTotal = InvalidRunningTotal + isInvalidIDPart1(idPair)
		} else {
			InvalidRunningTotal = InvalidRunningTotal + isInvalidIDPart2(idPair)
		}
	}
	return strconv.Itoa(InvalidRunningTotal)
}

func isInvalidIDPart1(idPair string) int {
	var totalInvalidIDs int = 0
	var splitIds []string = strings.Split(idPair, "-")
	var idStart, _ = strconv.Atoi(splitIds[0])
	var idEnd, _ = strconv.Atoi(splitIds[1])
	for i := idStart; i <= idEnd; i++ {
		var id string = strconv.Itoa(i)
		if utils.IsEven(len(id)) {
			var string1 string = id[0 : len(id)/2]
			var string2 string = id[len(id)/2:]
			if string1 == string2 {
				idNum, error := strconv.Atoi(id)
				if error != nil {
					log.Fatalf("Could not convert to integer: %v", error)
				}
				totalInvalidIDs = totalInvalidIDs + idNum
			}
		}
	}
	return totalInvalidIDs
}

func isInvalidIDPart2(idPair string) int {
	var totalInvalidIDs int = 0
	var splitIds []string = strings.Split(idPair, "-")
	var idStart, _ = strconv.Atoi(splitIds[0])
	var idEnd, _ = strconv.Atoi(splitIds[1])
	for i := idStart; i <= idEnd; i++ {
		var id string = strconv.Itoa(i)
		for j := 0; j < len(id); j++ {
			var string1 string = id[0:j]
			var string2 string = id[j:]
			string2 = strings.Replace(string2, string1, "", -1)
			if len(string2) == 0 {
				totalInvalidIDs = totalInvalidIDs + i
				break
			}

		}

	}
	return totalInvalidIDs
}
