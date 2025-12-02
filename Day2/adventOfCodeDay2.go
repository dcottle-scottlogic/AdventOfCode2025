package adventOfCodeDay2

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"AdventOfCode2025/utils"
)

func GetDay2Input() []string {
	var inputList = []string{"18623-26004", "226779-293422", "65855-88510", "868-1423", "248115026-248337139", "903911-926580", "97-121", "67636417-67796062", "24-47", "6968-10197", "193-242", "3769-5052", "5140337-5233474", "2894097247-2894150301", "979582-1016336", "502-646", "9132195-9191022", "266-378", "58-91", "736828-868857", "622792-694076", "6767592127-6767717303", "2920-3656", "8811329-8931031", "107384-147042", "941220-969217", "3-17", "360063-562672", "7979763615-7979843972", "1890-2660", "23170346-23308802"}
	return inputList
}

type Result struct {
	runningTotal int
}

func (r *Result) AddToTotal(value int) {
	r.runningTotal += value
}
func (r *Result) Part1HandleInvalidIDPair(idPair string) {
	var splitIds []string = strings.Split(idPair, "-")
	var idStart, _ = strconv.Atoi(splitIds[0])
	var idEnd, _ = strconv.Atoi(splitIds[1])
	for i := idStart; i <= idEnd; i++ {
		var id string = strconv.Itoa(i)
		if utils.IsEven(len(id)) {
			var string1 string = id[0 : len(id)/2]
			var string2 string = id[len(id)/2:]
			if string1 == string2 {
				idNum, _ := strconv.Atoi(id)
				r.AddToTotal(idNum)
			}
		}
	}
}
func (r *Result) Part2HandleInvalidIDPair(idPair string) {
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
				r.AddToTotal(i)
				break
			}

		}

	}
}
func (r *Result) GetTotal() int {
	return r.runningTotal
}

func Day2Main(isPart1 bool, input []string) string {
	newResult := Result{runningTotal: 0}
	fmt.Println("Starting Day 2 Computation")
	var wg sync.WaitGroup
	for i := range input {
		var idPair string = input[i]
		if isPart1 {
			wg.Go(func() { newResult.Part1HandleInvalidIDPair(idPair) })
		} else {
			wg.Go(func() { newResult.Part2HandleInvalidIDPair(idPair) })
		}
	}
	wg.Wait()
	return strconv.Itoa(newResult.GetTotal())
}
