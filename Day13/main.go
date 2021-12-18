package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/SotoDucani/AoC2021/internal/read"
)

func parseInput() (map[string][]int, [][]string) {
	input := read.ReadStrArrayByLine("./input.txt")

	coordMap := make(map[string][]int)
	var foldArray [][]string

	for _, line := range input {
		split := strings.Split(line, ",")
		if len(split) == 2 {
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])
			coord := []int{x, y}
			mapStr := fmt.Sprintf("%s,%s", split[0], split[1])
			coordMap[mapStr] = coord
		} else {
			foldSplit := strings.Split(line, " ")
			if len(foldSplit) == 3 {
				foldInfo := strings.Split(foldSplit[2], "=")
				fold := []string{foldInfo[0], foldInfo[1]}
				foldArray = append(foldArray, fold)
			}
		}
	}

	return coordMap, foldArray
}

func part1() {
	coordMap, foldArray := parseInput()

	// Only doing first fold, which is why this weird assignment is here
	// Otherwise we'd just loop on foldArray
	var fold [][]string
	fold = append(fold, foldArray[0])
	fmt.Printf("FoldInfo: %#v\n", fold)

	newCoordMap := make(map[string][]int)

	for _, foldInfo := range fold {
		foldInt, _ := strconv.Atoi(foldInfo[1])
		for coordStr, coord := range coordMap {
			if foldInfo[0] == "x" {
				if coord[0] > foldInt {
					distance := coord[0] - foldInt
					newCoord := foldInt - distance
					coord := []int{newCoord, coord[1]}
					mapStr := fmt.Sprintf("%s,%s", strconv.Itoa(newCoord), strconv.Itoa(coord[1]))
					newCoordMap[mapStr] = coord
				} else {
					newCoordMap[coordStr] = coord
				}
			} else if foldInfo[0] == "y" {
				if coord[1] > foldInt {
					distance := coord[1] - foldInt
					newCoord := foldInt - distance
					coord := []int{coord[0], newCoord}
					mapStr := fmt.Sprintf("%s,%s", strconv.Itoa(coord[0]), strconv.Itoa(newCoord))
					newCoordMap[mapStr] = coord
				} else {
					newCoordMap[coordStr] = coord
				}
			}
		}
	}

	//for _, coord := range newCoordMap {
	//	fmt.Printf("%#v\n", coord)
	//}

	fmt.Printf("Part 1 - Visible Dots: %d\n", len(newCoordMap))
}

func part2() {
	coordMap, foldArray := parseInput()

	var curCoordMap = coordMap

	for _, foldInfo := range foldArray {
		//fmt.Printf("FoldInfo: %#v\n", foldInfo)
		newCoordMap := make(map[string][]int)
		foldInt, _ := strconv.Atoi(foldInfo[1])
		for coordStr, coord := range curCoordMap {
			if foldInfo[0] == "x" {
				if coord[0] > foldInt {
					distance := coord[0] - foldInt
					newCoord := foldInt - distance
					coord := []int{newCoord, coord[1]}
					mapStr := fmt.Sprintf("%s,%s", strconv.Itoa(newCoord), strconv.Itoa(coord[1]))
					newCoordMap[mapStr] = coord
				} else {
					newCoordMap[coordStr] = coord
				}
			} else if foldInfo[0] == "y" {
				if coord[1] > foldInt {
					distance := coord[1] - foldInt
					newCoord := foldInt - distance
					coord := []int{coord[0], newCoord}
					mapStr := fmt.Sprintf("%s,%s", strconv.Itoa(coord[0]), strconv.Itoa(newCoord))
					newCoordMap[mapStr] = coord
				} else {
					newCoordMap[coordStr] = coord
				}
			}
		}
		curCoordMap = newCoordMap
	}

	//for _, coord := range curCoordMap {
	//	fmt.Printf("%#v\n", coord)
	//}

	// Find our max X and Y
	maxX := 0
	maxY := 0
	for _, v := range curCoordMap {
		if v[0] > maxX {
			maxX = v[0]
		}
		if v[1] > maxY {
			maxY = v[1]
		}
	}

	fmt.Printf("Part 2 - Code:\n\n")
	// Draw our resulting grid
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			mapStr := fmt.Sprintf("%s,%s", strconv.Itoa(x), strconv.Itoa(y))
			_, ok := curCoordMap[mapStr]
			if ok {
				fmt.Printf("X")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
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
