package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/SotoDucani/AoC2021/internal/read"
)

func parseFile() [][][]int {
	input := read.ReadStrArrayByLine("./input.txt")
	allLines := [][][]int{}

	for _, line := range input {
		largeSplit := strings.Split(line, " -> ")
		lineInfo := [][]int{}
		for _, pair := range largeSplit {
			smallSplit := strings.Split(pair, ",")
			x, _ := strconv.Atoi(smallSplit[0])
			y, _ := strconv.Atoi(smallSplit[1])
			coord := []int{x, y}
			lineInfo = append(lineInfo, coord)
		}
		allLines = append(allLines, lineInfo)
	}
	return allLines
}

func part1() {
	ventLines := parseFile()
	seaFloor := make(map[string]int)
	//fmt.Printf("%#v\n", ventLines)

	for _, line := range ventLines {
		if line[0][0] == line[1][0] {
			// Vertical line
			//fmt.Printf("Vert: %#v\n", line)
			var lower int
			var higher int
			// Determine lower/upper range for looping
			if line[0][1] < line[1][1] {
				lower = line[0][1]
				higher = line[1][1]
			} else {
				lower = line[1][1]
				higher = line[0][1]
			}

			// Create/Update our map
			for cur := lower; cur <= higher; cur++ {
				key := fmt.Sprintf("%v,%v", line[0][0], cur)
				//fmt.Printf("Vent found at: %#v\n", key)
				seaFloor[key] = seaFloor[key] + 1
			}
		} else if line[0][1] == line[1][1] {
			// Horizontal line
			//fmt.Printf("Hori: %#v\n", line)
			var lower int
			var higher int
			// Determine lower/upper range for looping
			if line[0][0] < line[1][0] {
				lower = line[0][0]
				higher = line[1][0]
			} else {
				lower = line[1][0]
				higher = line[0][0]
			}

			// Create/Update our map
			for cur := lower; cur <= higher; cur++ {
				key := fmt.Sprintf("%v,%v", cur, line[0][1])
				//fmt.Printf("Vent found at: %#v\n", key)
				seaFloor[key] = seaFloor[key] + 1
			}
		}
	}

	//fmt.Printf("seaFloor: %#v", seaFloor)

	var counter int
	for _, value := range seaFloor {
		if value >= 2 {
			//fmt.Printf("Location %#v has 2 or more vents\n", key)
			counter = counter + 1
		}
	}

	fmt.Printf("Part 1: %v\n", counter)
}

func part2() {
	ventLines := parseFile()
	seaFloor := make(map[string]int)
	//fmt.Printf("%#v\n", ventLines)

	for _, line := range ventLines {
		if line[0][0] == line[1][0] {
			// Vertical line
			//fmt.Printf("Vert: %#v\n", line)
			var lower int
			var higher int
			// Determine lower/upper range for looping
			if line[0][1] < line[1][1] {
				lower = line[0][1]
				higher = line[1][1]
			} else {
				lower = line[1][1]
				higher = line[0][1]
			}

			// Create/Update our map
			for cur := lower; cur <= higher; cur++ {
				key := fmt.Sprintf("%v,%v", line[0][0], cur)
				//fmt.Printf("Vent found at: %#v\n", key)
				seaFloor[key] = seaFloor[key] + 1
			}
		} else if line[0][1] == line[1][1] {
			// Horizontal line
			//fmt.Printf("Hori: %#v\n", line)
			var lower int
			var higher int
			// Determine lower/upper range for looping
			if line[0][0] < line[1][0] {
				lower = line[0][0]
				higher = line[1][0]
			} else {
				lower = line[1][0]
				higher = line[0][0]
			}

			// Create/Update our map
			for cur := lower; cur <= higher; cur++ {
				key := fmt.Sprintf("%v,%v", cur, line[0][1])
				//fmt.Printf("Vent found at: %#v\n", key)
				seaFloor[key] = seaFloor[key] + 1
			}
		} else {
			// Diagonal line
			//fmt.Printf("Diag: %#v\n", line)

			// Corrects our line info so we're always drawing towards the right, that is
			// we're moving from a lower x value to a higher one
			var correctedLineInfo [][]int
			if line[0][0] < line[1][0] {
				correctedLineInfo = line
			} else {
				correctedLineInfo = [][]int{line[1], line[0]}
			}

			if correctedLineInfo[0][1] < correctedLineInfo[1][1] {
				// This is a '\' diagonal
				x := correctedLineInfo[0][0]
				y := correctedLineInfo[0][1]
				for x <= correctedLineInfo[1][0] {
					key := fmt.Sprintf("%v,%v", x, y)
					//fmt.Printf("1Down diag Vent found at: %#v\n", key)
					seaFloor[key] = seaFloor[key] + 1
					x++
					y++
				}
			} else {
				// This is a '/' diagonal
				x := correctedLineInfo[0][0]
				y := correctedLineInfo[0][1]
				for x <= correctedLineInfo[1][0] {
					key := fmt.Sprintf("%v,%v", x, y)
					//fmt.Printf("2Down diag Vent found at: %#v\n", key)
					seaFloor[key] = seaFloor[key] + 1
					x++
					y--
				}
			}
		}
	}

	//fmt.Printf("seaFloor: %#v", seaFloor)

	var counter int
	for _, value := range seaFloor {
		if value >= 2 {
			//fmt.Printf("Location %#v has 2 or more vents\n", key)
			counter = counter + 1
		}
	}

	fmt.Printf("Part 2: %v", counter)
}

func main() {
	part1()
	part2()
}
