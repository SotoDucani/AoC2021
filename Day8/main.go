package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/SotoDucani/AoC2021/internal/read"
)

func parseInput() [][][]string {
	input := read.ReadStrArrayByLine("./input.txt")
	var parsed [][][]string
	for _, line := range input {
		sectionSplit := strings.Split(line, " | ")

		signals := strings.Split(sectionSplit[0], " ")
		outputs := strings.Split(sectionSplit[1], " ")
		linearray := [][]string{signals, outputs}
		parsed = append(parsed, linearray)
	}
	return parsed
}

func part1() {
	// 1,4,7,8
	input := parseInput()

	var count int
	for _, line := range input {
		for _, outputValue := range line[1] {
			//1
			if len(outputValue) == 2 {
				count = count + 1
			}
			//4
			if len(outputValue) == 4 {
				count = count + 1
			}
			//7
			if len(outputValue) == 3 {
				count = count + 1
			}
			//8
			if len(outputValue) == 7 {
				count = count + 1
			}
		}
	}

	fmt.Printf("Part 1 - Count: %v\n", count)
}

func findSharedSegments(knownNumber string, unknownNumber string) int {
	count := 0
	sharedMap := make(map[string]int)
	unknownEncodedSegments := read.StrToCharArray(unknownNumber)
	knownEncodedSegments := read.StrToCharArray(knownNumber)

	for _, char := range unknownEncodedSegments {
		sharedMap[char] = sharedMap[char] + 1
	}
	for _, char := range knownEncodedSegments {
		sharedMap[char] = sharedMap[char] + 1
		if sharedMap[char] == 2 {
			count = count + 1
		}
	}

	return count
}

func part2() {
	input := parseInput()

	resultSum := 0

	for _, signalLine := range input {
		//fmt.Printf("New line\n")
		// Known numbers
		knownNumber := make(map[string]string)

		// Map out known numbers
		for _, encodedNumber := range signalLine[0] {
			if len(encodedNumber) == 2 {
				//1
				knownNumber["1"] = encodedNumber
				//fmt.Printf("    1 known: %#v\n", encodedNumber)
			} else if len(encodedNumber) == 4 {
				//4
				knownNumber["4"] = encodedNumber
				//fmt.Printf("    4 known: %#v\n", encodedNumber)
			} else if len(encodedNumber) == 3 {
				//7
				knownNumber["7"] = encodedNumber
				//fmt.Printf("    7 known: %#v\n", encodedNumber)
			} else if len(encodedNumber) == 7 {
				//8
				knownNumber["8"] = encodedNumber
				//fmt.Printf("    8 known: %#v\n", encodedNumber)
			}
		}

		// Solve all other numbers by counting shared segments with known numbers to create a unique id
		// that identifies each 'unknown' number
		for _, encodedNumber := range signalLine[0] {
			// Skip known numbers
			if len(encodedNumber) == 2 || len(encodedNumber) == 4 || len(encodedNumber) == 3 || len(encodedNumber) == 7 {
				// Do nothing
			} else {
				// Shared segments with 1
				oneCount := findSharedSegments(knownNumber["1"], encodedNumber)

				// Shared segments with 4
				fourCount := findSharedSegments(knownNumber["4"], encodedNumber)

				// Shared segments with 7
				sevenCount := findSharedSegments(knownNumber["7"], encodedNumber)

				// Shared segments with 8
				eightCount := findSharedSegments(knownNumber["8"], encodedNumber)

				uniqueID := fmt.Sprintf("%v%v%v%v", oneCount, fourCount, sevenCount, eightCount)

				//fmt.Printf("Unique ID for %v: %v%v%v%v\n", encodedNumber, oneCount, fourCount, sevenCount, eightCount)

				// Based of some by-hand work to identify the unique number of shared segments with the known numbers
				if uniqueID == "2336" {
					// 0
					knownNumber["0"] = encodedNumber
					//fmt.Printf("    0 known: %#v\n", encodedNumber)
				} else if uniqueID == "1225" {
					// 2
					knownNumber["2"] = encodedNumber
					//fmt.Printf("    2 known: %#v\n", encodedNumber)
				} else if uniqueID == "2335" {
					// 3
					knownNumber["3"] = encodedNumber
					//fmt.Printf("    3 known: %#v\n", encodedNumber)
				} else if uniqueID == "1325" {
					// 5
					knownNumber["5"] = encodedNumber
					//fmt.Printf("    5 known: %#v\n", encodedNumber)
				} else if uniqueID == "1326" {
					// 6
					knownNumber["6"] = encodedNumber
					//fmt.Printf("    6 known: %#v\n", encodedNumber)
				} else if uniqueID == "2436" {
					// 9
					knownNumber["9"] = encodedNumber
					//fmt.Printf("    9 known: %#v\n", encodedNumber)
				}

			}
		}

		decodedOutput := ""
		for _, encodedOutput := range signalLine[1] {
			unknownEncodedSegments := read.StrToCharArray(encodedOutput)

			// Look for 100% segment match to a known number
			for k, knownNumberStr := range knownNumber {
				knownEncodedSegments := read.StrToCharArray(knownNumberStr)
				if len(knownEncodedSegments) == len(unknownEncodedSegments) {
					matchCount := findSharedSegments(knownNumberStr, encodedOutput)
					if matchCount == len(unknownEncodedSegments) {
						//fmt.Printf("    Output Segment %#v is a %v\n", encodedOutput, k)
						decodedOutput = decodedOutput + k
					}
				}
			}

		}
		intDecodedOutput, _ := strconv.Atoi(decodedOutput)
		resultSum = resultSum + intDecodedOutput
	}

	fmt.Printf("Part 2 - Result Sum: %v\n", resultSum)
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
