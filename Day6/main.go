package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/SotoDucani/AoC2021/internal/read"
)

func parseInput() []int {
	input := read.ReadStrArrayByLine("./input.txt")
	split := strings.Split(input[0], ",")
	var parsed []int
	for _, num := range split {
		convInt, _ := strconv.Atoi(num)
		parsed = append(parsed, convInt)
	}
	return parsed
}

func parseInputFor2() map[string]int {
	input := read.ReadStrArrayByLine("./input.txt")
	split := strings.Split(input[0], ",")
	var parsed []int
	for _, num := range split {
		convInt, _ := strconv.Atoi(num)
		parsed = append(parsed, convInt)
	}
	newMap := make(map[string]int)
	for _, value := range parsed {
		newMap[strconv.Itoa(value)] = newMap[strconv.Itoa(value)] + 1
	}
	return newMap
}

func part1() {
	parsed := parseInput()
	//fmt.Printf("Initial state: %#v\n", parsed)

	days := 80
	var newFish int
	for d := 1; d <= days; d++ {
		newFish = 0
		for i := 0; i < len(parsed); i++ {
			if parsed[i] == 0 {
				parsed[i] = 6
				newFish = newFish + 1
			} else {
				parsed[i] = parsed[i] - 1
			}
		}
		for f := 0; f < newFish; f++ {
			parsed = append(parsed, 8)
		}
		//fmt.Printf("After %v days: %#v\n", d, parsed)
	}

	fmt.Printf("Part 1 - Fish Count: %#v\n", len(parsed))
}

func part2() {
	parsed := parseInputFor2()
	//fmt.Printf("Initial state: %#v\n", parsed)

	days := 256
	for d := 1; d <= days; d++ {
		fishToReset := parsed["0"]
		parsed["0"] = parsed["1"]
		parsed["1"] = parsed["2"]
		parsed["2"] = parsed["3"]
		parsed["3"] = parsed["4"]
		parsed["4"] = parsed["5"]
		parsed["5"] = parsed["6"]
		parsed["6"] = parsed["7"] + fishToReset
		parsed["7"] = parsed["8"]
		parsed["8"] = fishToReset
	}

	var count int
	for _, v := range parsed {
		count = count + v
	}

	fmt.Printf("Part 2 - Fish Count: %v\n", count)
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
