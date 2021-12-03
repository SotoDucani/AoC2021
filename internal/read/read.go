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

func RemoveFromStringSlice(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func IntArrayToString(intarray []int) string {
	result := ""
	for _, num := range intarray {
		result = result + strconv.Itoa(num)
	}
	return result
}

func CharArrayToIntArray(str []string) []int {
	var ints []int
	for i := 0; i < len(str); i++ {
		cur, _ := strconv.Atoi(str[i])
		ints = append(ints, cur)
	}
	return ints
}

func IntToIntArray(number int) []int {
	numstr := strconv.Itoa(number)
	//fmt.Printf("%#v", numstr)
	strarray := strings.Split(numstr, "")
	//fmt.Printf("%#v", strarray)
	var intarray []int
	for _, i := range strarray {
		j, _ := strconv.Atoi(i)
		intarray = append(intarray, j)
	}
	return intarray
}

func StrToWordArray(str string) []string {
	words := strings.Split(str, " ")
	return words
}
