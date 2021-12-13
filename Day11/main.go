package main

import (
	"fmt"
	"time"

	"github.com/SotoDucani/AoC2021/internal/read"
)

type Octo struct {
	energy      int
	position    []int
	flashed     bool
	alreadyIncd bool
}

func parseInput() [][]Octo {
	input := read.ReadIntArrayByLine("./input.txt")

	var output [][]Octo
	yCount := 0
	for _, line := range input {
		lineArr := read.IntToIntArray(line)
		xCount := 0
		var octoLine []Octo
		for _, energyValue := range lineArr {
			curOcto := Octo{
				energy:      energyValue,
				position:    []int{xCount, yCount},
				flashed:     false,
				alreadyIncd: false,
			}
			octoLine = append(octoLine, curOcto)
			xCount++
		}
		output = append(output, octoLine)
		yCount++
	}
	return output
}

func flashOctoEnergy(targetOcto *Octo, octoGrid *[][]Octo) {
	curx := targetOcto.position[0]
	cury := targetOcto.position[1]
	lenx := len((*octoGrid)[cury]) //Manually dereference pointer
	leny := len((*octoGrid))       //Manually dereference pointer

	if !targetOcto.flashed {
		if targetOcto.energy > 9 {
			//fmt.Printf("Flashing Octo: %d,%d\n", curx, cury)
			targetOcto.flashed = true
			targetOcto.energy = 0

			incLeft, incRight, incUp, incDown := true, true, true, true
			//edge cases, lol, get it?
			//x
			switch curx {
			case 0:
				incLeft = false
			case lenx - 1:
				incRight = false
			}

			//y
			switch cury {
			case 0:
				incUp = false
			case leny - 1:
				incDown = false
			}

			if incLeft {
				//fmt.Printf("Incrementing %d,%d by %d,%d\n", curx-1, cury, curx, cury)
				if !(*octoGrid)[cury][curx-1].flashed {
					(*octoGrid)[cury][curx-1].energy++
				}
				flashOctoEnergy(&(*octoGrid)[cury][curx-1], octoGrid)
			}
			if incRight {
				//fmt.Printf("Incrementing %d,%d by %d,%d\n", curx+1, cury, curx, cury)
				if !(*octoGrid)[cury][curx+1].flashed {
					(*octoGrid)[cury][curx+1].energy++
				}
				flashOctoEnergy(&(*octoGrid)[cury][curx+1], octoGrid)
			}
			if incUp {
				//fmt.Printf("Incrementing %d,%d by %d,%d\n", curx, cury-1, curx, cury)
				if !(*octoGrid)[cury-1][curx].flashed {
					(*octoGrid)[cury-1][curx].energy++
				}
				flashOctoEnergy(&(*octoGrid)[cury-1][curx], octoGrid)
			}
			if incDown {
				//fmt.Printf("Incrementing %d,%d by %d,%d\n", curx, cury+1, curx, cury)
				if !(*octoGrid)[cury+1][curx].flashed {
					(*octoGrid)[cury+1][curx].energy++
				}
				flashOctoEnergy(&(*octoGrid)[cury+1][curx], octoGrid)
			}
			if incUp && incLeft {
				//fmt.Printf("Incrementing %d,%d by %d,%d\n", curx-1, cury-1, curx, cury)
				if !(*octoGrid)[cury-1][curx-1].flashed {
					(*octoGrid)[cury-1][curx-1].energy++
				}
				flashOctoEnergy(&(*octoGrid)[cury-1][curx-1], octoGrid)
			}
			if incUp && incRight {
				//fmt.Printf("Incrementing %d,%d by %d,%d\n", curx+1, cury-1, curx, cury)
				if !(*octoGrid)[cury-1][curx+1].flashed {
					(*octoGrid)[cury-1][curx+1].energy++
				}
				flashOctoEnergy(&(*octoGrid)[cury-1][curx+1], octoGrid)
			}
			if incDown && incLeft {
				//fmt.Printf("Incrementing %d,%d by %d,%d\n", curx-1, cury+1, curx, cury)
				if !(*octoGrid)[cury+1][curx-1].flashed {
					(*octoGrid)[cury+1][curx-1].energy++
				}
				flashOctoEnergy(&(*octoGrid)[cury+1][curx-1], octoGrid)
			}
			if incDown && incRight {
				//fmt.Printf("Incrementing %d,%d by %d,%d\n", curx+1, cury+1, curx, cury)
				if !(*octoGrid)[cury+1][curx+1].flashed {
					(*octoGrid)[cury+1][curx+1].energy++
				}
				flashOctoEnergy(&(*octoGrid)[cury+1][curx+1], octoGrid)
			}
		}
	}
}

func part1() {
	octoGrid := parseInput()

	flashCount := 0
	steps := 100

	for step := 1; step <= steps; step++ {
		//fmt.Printf("----- Starting Step %d -----\n", step)
		//Inc all octos first
		for y := 0; y < len(octoGrid); y++ {
			for x := 0; x < len(octoGrid[y]); x++ {
				octoGrid[y][x].energy++
				//fmt.Printf("Incrementing %d,%d to %d\n", x, y, octoGrid[y][x].energy)
			}
		}

		// Flash em
		for y := 0; y < len(octoGrid); y++ {
			for x := 0; x < len(octoGrid[y]); x++ {
				flashOctoEnergy(&octoGrid[y][x], &octoGrid)
			}
		}

		// Count and reset flashes
		for y := 0; y < len(octoGrid); y++ {
			for x := 0; x < len(octoGrid[y]); x++ {
				if octoGrid[y][x].flashed {
					//fmt.Printf("0")
					flashCount++
					octoGrid[y][x].flashed = false
				} else {
					//fmt.Printf("%d", octoGrid[y][x].energy)
				}
			}
			//fmt.Printf("\n")
		}
		//fmt.Printf("\n")
		//fmt.Printf("Current Flashes: %d\n", flashCount)
	}

	fmt.Printf("Part 1 - Flash count: %d\n", flashCount)
}

func part2() {
	octoGrid := parseInput()

	flashCount := 0
	step := 0
	steps := 500
	for step = 1; step <= steps && flashCount < 100; step++ {
		flashCount = 0
		fmt.Printf("----- Starting Step %d -----\n", step)
		//Inc all octos first
		for y := 0; y < len(octoGrid); y++ {
			for x := 0; x < len(octoGrid[y]); x++ {
				octoGrid[y][x].energy++
				//fmt.Printf("Incrementing %d,%d to %d\n", x, y, octoGrid[y][x].energy)
			}
		}

		// Flash em
		for y := 0; y < len(octoGrid); y++ {
			for x := 0; x < len(octoGrid[y]); x++ {
				flashOctoEnergy(&octoGrid[y][x], &octoGrid)
			}
		}

		// Count and reset flashes
		for y := 0; y < len(octoGrid); y++ {
			for x := 0; x < len(octoGrid[y]); x++ {
				if octoGrid[y][x].flashed {
					//fmt.Printf("0")
					flashCount++
					octoGrid[y][x].flashed = false
				} else {
					//fmt.Printf("%d", octoGrid[y][x].energy)
				}
			}
			//fmt.Printf("\n")
		}
		//fmt.Printf("\n")
		fmt.Printf("Step %d Flashes: %d\n", step-1, flashCount)
	}

	fmt.Printf("Part 2 - First Step for all flash: %d\n", step)
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
