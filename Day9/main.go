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
			if x == 0 {
				//fmt.Printf("testLeft false\n")
				testLeft = false
			}
			if x == len(input[0])-1 {
				//fmt.Printf("testRight false\n")
				testRight = false
			}
			if y == 0 {
				//fmt.Printf("testUp false\n")
				testUp = false
			}
			if y == len(input)-1 {
				//fmt.Printf("testDown false\n")
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

func searchBasin(curLocation []int) [][]int {
	input := parseInput()

	var basinMembers [][]int

	x := curLocation[0]
	y := curLocation[1]
	cur := input[y][x]

	//fmt.Printf("Checking Spot @ (x,y): (%d,%d)\n", x, y)

	testUp, testDown, testRight, testLeft := true, true, true, true
	//edge cases, lol, get it?
	if x == 0 {
		//fmt.Printf("testLeft false\n")
		testLeft = false
	}
	if x == len(input[0])-1 {
		//fmt.Printf("testRight false\n")
		testRight = false
	}
	if y == 0 {
		//fmt.Printf("testUp false\n")
		testUp = false
	}
	if y == len(input)-1 {
		//fmt.Printf("testDown false\n")
		testDown = false
	}

	if testUp {
		//fmt.Printf("testUp\n")
		check := input[y-1][x]
		if check > cur && check != 9 {
			foundBasinMembers := searchBasin([]int{x, y - 1})
			basinMembers = append(basinMembers, foundBasinMembers...)
		}
	}

	if testDown {
		//fmt.Printf("testDown\n")
		check := input[y+1][x]
		if check > cur && check != 9 {
			foundBasinMembers := searchBasin([]int{x, y + 1})
			basinMembers = append(basinMembers, foundBasinMembers...)
		}
	}

	if testLeft {
		//fmt.Printf("testLeft\n")
		check := input[y][x-1]
		if check > cur && check != 9 {
			foundBasinMembers := searchBasin([]int{x - 1, y})
			basinMembers = append(basinMembers, foundBasinMembers...)
		}
	}

	if testRight {
		//fmt.Printf("testRight\n")
		check := input[y][x+1]
		if check > cur && check != 9 {
			foundBasinMembers := searchBasin([]int{x + 1, y})
			basinMembers = append(basinMembers, foundBasinMembers...)
		}
	}

	//fmt.Printf("Appending (%d,%d) to found\n", x, y)
	basinMembers = append(basinMembers, []int{x, y})

	return basinMembers
}

func part2(lowSpots [][]int) {
	var basinSizes []int

	for _, lowSpot := range lowSpots {
		basinMap := make(map[string]int)

		foundMembers := searchBasin(lowSpot)
		//fmt.Printf("For basin (%d,%d) found members:%#v\n", lowSpot[0], lowSpot[1], foundMembers)

		for _, v := range foundMembers {
			locStr := fmt.Sprintf("%d,%d", v[0], v[1])
			basinMap[locStr] = basinMap[locStr] + 1
		}

		basinSize := len(basinMap)
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
