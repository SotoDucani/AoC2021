package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/SotoDucani/AoC2021/internal/read"
)

func parseInput() [][]string {
	input := read.ReadStrArrayByLine("./input.txt")

	var output [][]string
	var characters []string
	for _, line := range input {
		characters = read.StrToCharArray(line)
		output = append(output, characters)
	}
	return output
}

func part1() {
	input := parseInput()

	result := 0

	scoreMap := make(map[string]int)
	scoreMap[")"] = 3
	scoreMap["]"] = 57
	scoreMap["}"] = 1197
	scoreMap[">"] = 25137

	expectedMap := make(map[string]string)
	expectedMap["["] = "]"
	expectedMap["{"] = "}"
	expectedMap["<"] = ">"
	expectedMap["("] = ")"

	expectedMapkeys := make([]string, 0, len(expectedMap))
	for k := range expectedMap {
		expectedMapkeys = append(expectedMapkeys, k)
	}

	currentLine := 0
	for _, line := range input {
		var parsedStack read.Stack
		var expectedStack read.Stack
		currentCharacter := 0
		for _, v := range line {
			if read.SliceContains(expectedMapkeys, v) {
				// Opening bracket
				parsedStack.Push(v)
				expectedStack.Push(expectedMap[v])
			} else {
				// Closing bracket
				_, parsedPresent := parsedStack.Pop()
				expectedPop, expectedPresent := expectedStack.Pop()

				if parsedPresent && expectedPresent {
					if expectedPop == v {
						//do nothin
					} else {
						//fmt.Printf("Mismatched bracket at char %d, line %d: Expected %s, Got %s\n", currentCharacter, currentLine, expectedPop, v)
						result = result + scoreMap[v]
						break
					}
				} else if !parsedPresent {
					fmt.Printf("Ran out of parsedStack %d\n", currentCharacter)
					break
				} else if !expectedPresent {
					fmt.Printf("Ran out of expectedStack  %d\n", currentCharacter)
					break
				}
			}
			currentCharacter = currentCharacter + 1
		}
		currentLine = currentLine + 1
	}

	fmt.Printf("Part 1 - Score: %d\n", result)
}

func part2() {
	input := parseInput()

	var lineScores []int

	scoreMap := make(map[string]int)
	scoreMap[")"] = 1
	scoreMap["]"] = 2
	scoreMap["}"] = 3
	scoreMap[">"] = 4

	expectedMap := make(map[string]string)
	expectedMap["["] = "]"
	expectedMap["{"] = "}"
	expectedMap["<"] = ">"
	expectedMap["("] = ")"

	expectedMapkeys := make([]string, 0, len(expectedMap))
	for k := range expectedMap {
		expectedMapkeys = append(expectedMapkeys, k)
	}

	currentLine := 0
	for _, line := range input {
		var parsedStack read.Stack
		var expectedStack read.Stack
		currentCharacter := 0
		lineScore := 0
		calculateScore := true
		for _, v := range line {
			if read.SliceContains(expectedMapkeys, v) {
				// Opening bracket
				parsedStack.Push(v)
				expectedStack.Push(expectedMap[v])
			} else {
				// Closing bracket
				_, parsedPresent := parsedStack.Pop()
				expectedPop, expectedPresent := expectedStack.Pop()

				if parsedPresent && expectedPresent {
					if expectedPop == v {
						//do nothin
					} else {
						//fmt.Printf("Mismatched bracket at char %d, line %d: Expected %s, Got %s\n", currentCharacter, currentLine, expectedPop, v)
						//result = result + scoreMap[v]
						calculateScore = false
					}
				} else if !parsedPresent {
					fmt.Printf("Ran out of parsedStack %d\n", currentCharacter)
				} else if !expectedPresent {
					fmt.Printf("Ran out of expectedStack  %d\n", currentCharacter)
				}
			}
			currentCharacter = currentCharacter + 1
		}
		if calculateScore {
			//fmt.Printf("Line %#v", line)
			//fmt.Printf("Remaining expected: %#v\n", expectedStack)
			// Backwards loop because I used a 'stack'
			for i := len(expectedStack) - 1; i > -1; i-- {
				lineScore = (lineScore * 5) + scoreMap[expectedStack[i]]
			}
			lineScores = append(lineScores, lineScore)
		}
		currentLine = currentLine + 1
	}
	sort.Ints(lineScores)
	result := lineScores[(len(lineScores)-1)/2]

	fmt.Printf("Part 2 - Score: %#v\n", result)
}

func main() {
	p1b := time.Now()
	part1()
	mid := time.Now()
	part2()
	p2a := time.Now()
	part1Time := mid.Sub(p1b)
	part2Time := p2a.Sub(mid)
	fmt.Printf("Part 1 Time: %vμs\nPart 2 Time: %vμs\n", part1Time.Microseconds(), part2Time.Microseconds())
}
