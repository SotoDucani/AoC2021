package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/SotoDucani/AoC2021/internal/read"
)

func parseInput() ([]string, map[string]string) {
	input := read.ReadStrArrayByLine("./input.txt")

	var template []string
	rules := make(map[string]string)

	for _, line := range input {
		pairSplit := strings.Split(line, " -> ")
		if len(pairSplit) == 2 {
			rules[pairSplit[0]] = pairSplit[1]
		} else if line == "" {
			//donothing
		} else {
			template = read.StrToCharArray(line)
		}
	}

	return template, rules
}

func part1() {
	template, rules := parseInput()

	steps := 10

	curPoly := template

	for s := 0; s < steps; s++ {
		// Parse out pairs
		var pairs []string
		for charPos := 0; charPos < len(curPoly)-1; charPos++ {
			pairStr := curPoly[charPos] + curPoly[charPos+1]
			pairs = append(pairs, pairStr)
		}

		// Inject using pair insersion rules
		curPoly = nil
		for _, pair := range pairs {
			//fmt.Printf("Pair: %s\n", pair)
			//fmt.Printf("Inserting: %s\n", rules[pair])
			pairChars := read.StrToCharArray(pair)
			curPoly = append(curPoly, pairChars[0])
			curPoly = append(curPoly, rules[pair])
		}
		lastChar := read.StrToCharArray(pairs[len(pairs)-1])[1]
		curPoly = append(curPoly, lastChar)
		//fmt.Printf("After step %d: %#v\n", s+1, curPoly)
	}

	//Troubleshoot part 2 - step count must match between part1 and part2 to use this
	/*
		pairMap := make(map[string]int)
		for charPos := 0; charPos < len(curPoly)-1; charPos++ {
			pairStr := curPoly[charPos] + curPoly[charPos+1]
			pairMap[pairStr]++
		}
		fmt.Println(pairMap)
	*/

	// Find quantities
	quantMap := make(map[string]int)
	for _, char := range curPoly {
		quantMap[char]++
	}

	//fmt.Println(quantMap)
	var least int
	var most int
	for _, v := range quantMap {
		if least == 0 {
			least = v
		} else if v < least {
			least = v
		}

		if most == 0 {
			most = v
		} else if v > most {
			most = v
		}
	}

	fmt.Printf("Part 1 - Calc: %d\n", most-least)
}

func part2() {
	template, rules := parseInput()

	steps := 40

	pairMap := make(map[string]int)
	for charPos := 0; charPos < len(template)-1; charPos++ {
		pairStr := template[charPos] + template[charPos+1]
		pairMap[pairStr]++
	}

	//fmt.Println("Initial:", pairMap)

	for s := 0; s < steps; s++ {
		// Inject using pair insersion rules
		update := make(map[string]int)
		for pair, v := range pairMap {
			//fmt.Printf("Pair: %s\n", pair)
			//fmt.Printf("Inserting: %s\n", rules[pair])
			pairChars := read.StrToCharArray(pair)
			newChar := rules[pair]
			pair1 := pairChars[0] + newChar
			pair2 := newChar + pairChars[1]
			update[pair1] += v
			update[pair2] += v
		}
		pairMap = update
		//fmt.Printf("After step %d:\n", s+1)
		//fmt.Println("    ", pairMap)
		//fmt.Printf("After step %d\n", s+1)
	}

	//fmt.Println(pairMap)
	counts := make(map[string]int)
	for k, v := range pairMap {
		counts[string(k[0])] += v
	}
	// Add last character from template
	counts[string(template[len(template)-1])]++

	//fmt.Println(counts)
	var least int
	var most int
	for _, v := range counts {
		if least == 0 {
			least = v
		} else if v < least {
			least = v
		}

		if most == 0 {
			most = v
		} else if v > most {
			most = v
		}
	}

	fmt.Printf("Part 2 - Calc: %d\n", most-least)
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
