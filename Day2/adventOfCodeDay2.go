package adventOfCodeDay2

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"AdventOfCode2025/utils"
)

func GetDay2Input() []string {
	var inputList = []string{}
	return inputList
}

type Day2Result struct {
	runningTotal int
}

func (r *Day2Result) AddToTotal(value int) {
	r.runningTotal += value
}
func (r *Day2Result) Part1HandleInvalidIDPair(idPair string) {
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
func (r *Day2Result) Part2HandleInvalidIDPair(idPair string) {
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
func (r *Day2Result) GetTotal() int {
	return r.runningTotal
}

func Day2Main(isPart1 bool, input []string) string {
	newResult := Day2Result{runningTotal: 0}
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
