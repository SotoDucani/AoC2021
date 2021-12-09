package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	read "github.com/SotoDucani/AoC2021/internal/read"
)

func part1() {
	instructions := read.ReadStrArrayByLine("./input.txt")
	horizontal := 0
	depth := 0

	for _, stepStr := range instructions {
		stepSli := read.StrToWordArray(stepStr)
		switch stepSli[0] {
		case "forward":
			amt, err := strconv.Atoi(stepSli[1])
			if err != nil {
				log.Fatalf("Failed in forward: %v\n", err)
			}
			horizontal = horizontal + amt
		case "up":
			amt, err := strconv.Atoi(stepSli[1])
			if err != nil {
				log.Fatalf("Failed in up: %v\n", err)
			}
			depth = depth - amt
		case "down":
			amt, err := strconv.Atoi(stepSli[1])
			if err != nil {
				log.Fatalf("Failed in down: %v\n", err)
			}
			depth = depth + amt
		}

	}
	result := horizontal * depth
	fmt.Printf("Horizontal: %v Depth: %v Result: %v\n", horizontal, depth, result)
}

func part2() {
	instructions := read.ReadStrArrayByLine("./input.txt")
	horizontal := 0
	depth := 0
	aim := 0

	for _, stepStr := range instructions {
		stepSli := read.StrToWordArray(stepStr)
		switch stepSli[0] {
		case "forward":
			amt, err := strconv.Atoi(stepSli[1])
			if err != nil {
				log.Fatalf("Failed in forward: %v\n", err)
			}
			horizontal = horizontal + amt
			depth = depth + (amt * aim)
		case "up":
			amt, err := strconv.Atoi(stepSli[1])
			if err != nil {
				log.Fatalf("Failed in up: %v\n", err)
			}
			aim = aim - amt
			//.Printf("Aim now: %v", aim)
		case "down":
			amt, err := strconv.Atoi(stepSli[1])
			if err != nil {
				log.Fatalf("Failed in down: %v\n", err)
			}
			aim = aim + amt
			//log.Printf("Aim now: %v", aim)
		}

	}
	result := horizontal * depth
	fmt.Printf("Horizontal: %v Depth: %v Aim: %v Result: %v\n", horizontal, depth, aim, result)
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
