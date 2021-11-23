package read

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadIntArrayByLine(filepath string) []int {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var intarray []int
	for scanner.Scan() {
		cur, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		intarray = append(intarray, cur)
	}

	return intarray
}

func ReadStrArrayByLine(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var strarray []string
	for scanner.Scan() {
		cur := scanner.Text()
		strarray = append(strarray, cur)
	}

	return strarray
}

func StrToCharArray(str string) []string {
	runes := []rune(str)
	var chars []string
	for i := 0; i < len(runes); i++ {
		chars = append(chars, string(runes[i]))
	}
	return chars
}

func StrToWordArray(str string) []string {
	words := strings.Split(str, " ")
	return words
}
