package main

import (
	"fmt"

	read "github.com/SotoDucani/AoC2021/internal/read"
)

func part1() {
	array := read.ReadIntArrayByLine("./input.txt")
	for x := 0; x < len(array); x++ {
		value1 := array[x]
		for y := (x + 1); y < len(array); y++ {
			value2 := array[y]
			if (value1 + value2) == 2020 {
				fmt.Printf("Part 1: %d\n", (value1 * value2))
			}
		}
	}
}

func part2() {
	array := read.ReadIntArrayByLine("./input.txt")
	for x := 0; x < len(array); x++ {
		value1 := array[x]
		for y := (x + 1); y < len(array); y++ {
			value2 := array[y]
			for z := (x + 2); z < len(array); z++ {
				value3 := array[z]
				if (value1 + value2 + value3) == 2020 {
					fmt.Printf("Part 2: %d\n", (value1 * value2 * value3))
				}
			}
		}
	}
}

func main() {
	part1()
	part2()
}
