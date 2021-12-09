package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/SotoDucani/AoC2021/internal/read"
)

func part1() {
	gamma := make([]int, 12)
	epsilon := make([]int, 12)
	count := make([]int, 12)

	diag := read.ReadStrArrayByLine("./input.txt")
	//fmt.Printf("%#v", count)
	for _, line := range diag {
		chararray := read.StrToCharArray(line)
		//fmt.Printf("%#v", chararray)
		intarray := read.CharArrayToIntArray(chararray)
		count[0] = count[0] + intarray[0]
		count[1] = count[1] + intarray[1]
		count[2] = count[2] + intarray[2]
		count[3] = count[3] + intarray[3]
		count[4] = count[4] + intarray[4]
		count[5] = count[5] + intarray[5]
		count[6] = count[6] + intarray[6]
		count[7] = count[7] + intarray[7]
		count[8] = count[8] + intarray[8]
		count[9] = count[9] + intarray[9]
		count[10] = count[10] + intarray[10]
		count[11] = count[11] + intarray[11]
	}

	for i := 0; i < len(count); i++ {
		if count[i] <= (len(diag) / 2) {
			//fmt.Printf("Count: %v Value: %v Less than %v\n", i, count[i], len(diag)/2)
			gamma[i] = 0
			epsilon[i] = 1
		} else {
			//fmt.Printf("Count: %v Value: %v Greater than %v\n", i, count[i], len(diag)/2)
			gamma[i] = 1
			epsilon[i] = 0
		}
	}

	strGamma := read.IntArrayToString(gamma)
	strEpsilon := read.IntArrayToString(epsilon)

	b64Gamma, _ := strconv.ParseInt(strGamma, 2, 64)
	b64Epsilon, _ := strconv.ParseInt(strEpsilon, 2, 64)
	//fmt.Printf("Count: %#v\n", count)
	//fmt.Printf("Gamma: %#v\n", b64Gamma)
	//fmt.Printf("Epsilon:%#v\n", b64Epsilon)
	fmt.Printf("Part 1 Result: %#v\n", b64Gamma*b64Epsilon)
}

func part2() {
	file := read.ReadStrArrayByLine("./input.txt")
	//Oxy
	diag := file
	for i := 0; i < 12; i++ {
		var zeroArray []string
		var oneArray []string

		for _, line := range diag {
			charArray := read.StrToCharArray(line)
			if charArray[i] == "0" {
				zeroArray = append(zeroArray, line)
			} else {
				oneArray = append(oneArray, line)
			}
		}
		//fmt.Printf("ZeroLen: %#v\n", len(zeroArray))
		//fmt.Printf("OneLen:%#v\n", len(oneArray))
		if (len(zeroArray) + len(oneArray)) == 1 {
			break
		}
		if len(zeroArray) > len(oneArray) {
			diag = zeroArray
			//fmt.Printf("Keeping zeroArray\n")
		} else {
			diag = oneArray
			//fmt.Printf("Keeping oneArray\n")
		}
	}
	oxyStrArray := diag
	//fmt.Printf("%#v\n", oxyStrArray)

	//CO2
	diag = file
	for i := 0; i < 12; i++ {
		var zeroArray []string
		var oneArray []string

		for _, line := range diag {
			charArray := read.StrToCharArray(line)
			if charArray[i] == "0" {
				zeroArray = append(zeroArray, line)
			} else {
				oneArray = append(oneArray, line)
			}
		}
		//fmt.Printf("ZeroLen: %#v\n", len(zeroArray))
		//fmt.Printf("OneLen:%#v\n", len(oneArray))
		if (len(zeroArray) + len(oneArray)) == 1 {
			break
		}
		if len(zeroArray) > len(oneArray) {
			diag = oneArray
			//fmt.Printf("Keeping oneArray\n")
		} else {
			diag = zeroArray
			//fmt.Printf("Keeping zeroArray\n")
		}
	}
	co2StrArray := diag
	//fmt.Printf("%#v\n", co2StrArray)

	b64oxyStrArray, _ := strconv.ParseInt(oxyStrArray[0], 2, 64)
	b64co2StrArray, _ := strconv.ParseInt(co2StrArray[0], 2, 64)

	fmt.Printf("Part 2 Result: %#v\n", b64co2StrArray*b64oxyStrArray)
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
