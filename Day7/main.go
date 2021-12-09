package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/SotoDucani/AoC2021/internal/read"
)

func parseInput() (map[string]int, int, int) {
	input := read.ReadStrArrayByLine("./input.txt")
	split := strings.Split(input[0], ",")
	var sortSlice []int
	var parsed []int
	for _, num := range split {
		convInt, _ := strconv.Atoi(num)
		parsed = append(parsed, convInt)
		sortSlice = append(sortSlice, convInt)
	}
	newMap := make(map[string]int)
	for _, value := range parsed {
		newMap[strconv.Itoa(value)] = newMap[strconv.Itoa(value)] + 1
	}

	sort.Ints(sortSlice)
	return newMap, sortSlice[0], sortSlice[len(sortSlice)-1]
}

func part1() {
	input, min, max := parseInput()
	//fmt.Printf("%#v,%v,%v", input, min, max)

	positionCost := make(map[int]int)

	for cur := min; cur <= max; cur++ {
		for k, v := range input {
			numK, _ := strconv.Atoi(k)
			distance := int(math.Abs(float64(numK - cur)))
			curCost := distance * v
			positionCost[cur] = positionCost[cur] + curCost
		}
	}

	var values []int
	for _, v := range positionCost {
		values = append(values, v)
	}
	sort.Ints(values)

	lowestValue := values[0]
	lowestPosition := 0

	for k, v := range positionCost {
		//fmt.Printf("Key: %v Value: %v\n", k, v)
		if v == lowestValue {
			lowestPosition = k
		}
	}

	fmt.Printf("Part 1 - Position: %v Fuel Total: %v\n", lowestPosition, lowestValue)
}

func part2() {
	input, min, max := parseInput()
	//fmt.Printf("%#v,%v,%v", input, min, max)

	positionCost := make(map[int]int)

	for cur := min; cur <= max; cur++ {
		for k, v := range input {
			numK, _ := strconv.Atoi(k)
			distance := int(math.Abs(float64(numK - cur)))
			sumCost := 0
			for stepCost := 1; stepCost <= distance; stepCost++ {
				sumCost = sumCost + stepCost
			}
			positionCost[cur] = positionCost[cur] + (sumCost * v)

		}
	}

	var values []int
	for _, v := range positionCost {
		values = append(values, v)
	}
	sort.Ints(values)

	lowestValue := values[0]
	lowestPosition := 0

	for k, v := range positionCost {
		//fmt.Printf("Key: %v Value: %v\n", k, v)
		if v == lowestValue {
			lowestPosition = k
		}
	}

	fmt.Printf("Part 2 - Position: %v Fuel Total: %v\n", lowestPosition, lowestValue)
}

func main() {
	p1b := time.Now()
	part1()
	mid := time.Now()
	part2()
	p2a := time.Now()
	part1Time := mid.Sub(p1b)
	part2Time := p2a.Sub(mid)
	fmt.Printf("Part 1 Time: %dμs\nPart 2 Time: %dμs\n", part1Time.Microseconds(), part2Time.Microseconds())
}
