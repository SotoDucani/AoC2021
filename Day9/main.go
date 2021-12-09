package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/SotoDucani/AoC2021/internal/read"
)

func parseInput() [][]int {
	input := read.ReadStrArrayByLine("./input.txt")
	var output [][]int
	for _, line := range input {
		var lineInt []int
		splitLine := read.StrToCharArray(line)
		for _, char := range splitLine {
			num, _ := strconv.Atoi(char)
			lineInt = append(lineInt, num)
		}
		output = append(output, lineInt)
	}
	return output
}

func part1() [][]int {
	input := parseInput()

	var riskLevel int
	var lowSpots [][]int

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			cur := input[y][x]
			//fmt.Printf("Current: %d\n", cur)
			testUp, testDown, testRight, testLeft := true, true, true, true

			//edge cases, lol, get it?
			switch x {
			case 0:
				testLeft = false
			case len(input[0]) - 1:
				testRight = false
			}

			switch y {
			case 0:
				testUp = false
			case len(input) - 1:
				testDown = false
			}

			lowSpot := true
			//Testing
			if testUp {
				//fmt.Printf("Checking Spot @ (x,y): (%d,%d)\n", x, y)
				check := input[y-1][x]
				if check <= cur {
					lowSpot = false
				}
			}

			if testDown {
				check := input[y+1][x]
				if check <= cur {
					lowSpot = false
				}
			}

			if testLeft {
				check := input[y][x-1]
				if check <= cur {
					lowSpot = false
				}
			}

			if testRight {
				check := input[y][x+1]
				if check <= cur {
					lowSpot = false
				}
			}

			//If we're still a low spot
			if lowSpot {
				//fmt.Printf("Low Spot @ (x,y): (%d,%d)\n", x, y)
				lowSpots = append(lowSpots, []int{x, y})
				riskLevel = riskLevel + (1 + input[y][x])
			}
		}
	}

	fmt.Printf("Part 1 - Risk level: %d\n", riskLevel)
	return lowSpots
}

func searchBasin(input [][]int, curLocation []int, searchedLocations map[string]int) [][]int {

	var basinMembers [][]int

	x := curLocation[0]
	y := curLocation[1]
	cur := input[y][x]
	mapString := fmt.Sprintf("%d,%c", x, y)
	searchedLocations[mapString] = searchedLocations[mapString] + 1

	//fmt.Printf("Checking Spot @ (x,y): (%d,%d)\n", x, y)

	if searchedLocations[mapString] == 1 {
		testUp, testDown, testRight, testLeft := true, true, true, true
		//edge cases, lol, get it?
		switch x {
		case 0:
			testLeft = false
		case len(input[0]) - 1:
			testRight = false
		}

		switch y {
		case 0:
			testUp = false
		case len(input) - 1:
			testDown = false
		}

		if testUp {
			//fmt.Printf("testUp\n")
			check := input[y-1][x]
			if check > cur && check != 9 {
				foundBasinMembers := searchBasin(input, []int{x, y - 1}, searchedLocations)
				basinMembers = append(basinMembers, foundBasinMembers...)
			}
		}

		if testDown {
			//fmt.Printf("testDown\n")
			check := input[y+1][x]
			if check > cur && check != 9 {
				foundBasinMembers := searchBasin(input, []int{x, y + 1}, searchedLocations)
				basinMembers = append(basinMembers, foundBasinMembers...)
			}
		}

		if testLeft {
			//fmt.Printf("testLeft\n")
			check := input[y][x-1]
			if check > cur && check != 9 {
				foundBasinMembers := searchBasin(input, []int{x - 1, y}, searchedLocations)
				basinMembers = append(basinMembers, foundBasinMembers...)
			}
		}

		if testRight {
			//fmt.Printf("testRight\n")
			check := input[y][x+1]
			if check > cur && check != 9 {
				foundBasinMembers := searchBasin(input, []int{x + 1, y}, searchedLocations)
				basinMembers = append(basinMembers, foundBasinMembers...)
			}
		}

		//fmt.Printf("Appending (%d,%d) to found\n", x, y)
		basinMembers = append(basinMembers, []int{x, y})
	}

	return basinMembers
}

func part2(lowSpots [][]int) {
	input := parseInput()
	var basinSizes []int

	for _, lowSpot := range lowSpots {
		searchedLocations := make(map[string]int)

		foundMembers := searchBasin(input, lowSpot, searchedLocations)
		//fmt.Printf("For basin (%d,%d) found members:%#v\n", lowSpot[0], lowSpot[1], foundMembers)

		basinSize := len(foundMembers)
		//fmt.Printf("Basin Size found: %d\n", basinSize)
		basinSizes = append(basinSizes, basinSize)
	}

	//fmt.Printf("BasinSizes: %#v\n", basinSizes)

	first := 0
	second := 0
	third := 0

	for _, size := range basinSizes {
		if size > first {
			third = second
			second = first
			first = size
		} else if size > second {
			third = second
			second = size
		} else if size > third {
			third = size
		}
	}

	result := first * second * third

	fmt.Printf("Part 2 - 3 Largest Basin multi: %d\n", result)
}

func main() {
	p1b := time.Now()
	lowSpots := part1()
	mid := time.Now()
	part2(lowSpots)
	p2a := time.Now()
	part1Time := mid.Sub(p1b)
	part2Time := p2a.Sub(mid)
	fmt.Printf("Part 1 Time: %vμs\nPart 2 Time: %vμs\n", part1Time.Microseconds(), part2Time.Microseconds())
}
