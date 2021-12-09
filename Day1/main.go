package main

import (
	"fmt"
	"time"

	read "github.com/SotoDucani/AoC2021/internal/read"
)

func part1() {
	array := read.ReadIntArrayByLine("./input.txt")
	var (
		increaseCount int
		decreaseCount int
		sameCount     int
	)
	for x := 1; x < len(array); x++ {
		prev := array[x-1]
		cur := array[x]
		if prev < cur {
			increaseCount = increaseCount + 1
		} else if cur < prev {
			decreaseCount = decreaseCount + 1
		} else {
			sameCount = sameCount + 1
		}
	}
	fmt.Printf("Part 1 - Increase: %v Decrease: %v Same: %v\n", increaseCount, decreaseCount, sameCount)
}

func part2() {
	array := read.ReadIntArrayByLine("./input.txt")
	var (
		increaseCount int
		decreaseCount int
		sameCount     int
	)
	for x := 3; x < len(array); x++ {
		prevA := array[x-3]
		prevB := array[x-2]
		prevC := array[x-1]
		prevSum := prevA + prevB + prevC

		curA := array[x-2]
		curB := array[x-1]
		curC := array[x]
		curSum := curA + curB + curC
		if prevSum < curSum {
			increaseCount = increaseCount + 1
		} else if curSum < prevSum {
			decreaseCount = decreaseCount + 1
		} else {
			sameCount = sameCount + 1
		}
	}
	fmt.Printf("Part 2 - Increase: %v Decrease: %v Same: %v\n", increaseCount, decreaseCount, sameCount)
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
