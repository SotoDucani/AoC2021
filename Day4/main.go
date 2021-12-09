package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/SotoDucani/AoC2021/internal/read"
)

type Cell struct {
	number int
	marked bool
}

type BingoBoard struct {
	board        [][]Cell
	winningBoard bool
	winningNum   int
	stepsToWin   int
	unmarkedSum  int
}

func parseFile() ([]int, []BingoBoard) {
	input := read.ReadStrArrayByLine("./input.txt")

	//First Line
	var calledNumbers []int
	split := strings.Split(input[0], ",")
	for _, str := range split {
		num, _ := strconv.Atoi(str)
		calledNumbers = append(calledNumbers, num)
	}

	//Matrices
	var allMatrices []BingoBoard
	var matrix [][]Cell
	for i := 2; i < len(input); i++ {
		//fmt.Printf("Parsing Line: %#v\n", input[i])

		// Skip blank lines
		if input[i] == "" {
			//fmt.Printf("Blank Line\n")
			//fmt.Printf("Matrix : %#v\n", matrix)
			new := BingoBoard{
				board:       matrix,
				stepsToWin:  0,
				unmarkedSum: 0,
			}
			allMatrices = append(allMatrices, new)
			matrix = nil
			continue
		}

		// Iterate the row
		var rowArray []Cell
		lineSplit := strings.Split(input[i], " ")
		//fmt.Printf("Line Split: %#v\n", lineSplit)
		for _, str := range lineSplit {
			num, _ := strconv.Atoi(str)
			curCell := Cell{
				number: num,
			}
			rowArray = append(rowArray, curCell)
		}
		//fmt.Printf("RowArray : %#v\n", rowArray)
		matrix = append(matrix, rowArray)
	}
	// Append last
	new := BingoBoard{
		board: matrix,
	}
	allMatrices = append(allMatrices, new)
	return calledNumbers, allMatrices
}

func checkWin(curBoard *BingoBoard) {
	// Check for row wins
	for rowPos := 0; rowPos < len(curBoard.board); rowPos++ {
		markedInRow := 0
		for colPos := 0; colPos < len(curBoard.board[rowPos]); colPos++ {
			if curBoard.board[rowPos][colPos].marked {
				markedInRow = markedInRow + 1
				if markedInRow == len(curBoard.board[rowPos]) {
					//fmt.Printf("    Row Win detected, %#v steps\n", curBoard.stepsToWin)
					curBoard.winningBoard = true
				}
			}
		}
	}
	// Check for col wins
	for colPos := 0; colPos < len(curBoard.board[0]); colPos++ {
		markedInCol := 0
		for rowPos := 0; rowPos < len(curBoard.board); rowPos++ {
			if curBoard.board[rowPos][colPos].marked {
				markedInCol = markedInCol + 1
				if markedInCol == len(curBoard.board[rowPos]) {
					//fmt.Printf("    Col Win detected, %#v steps\n", curBoard.stepsToWin)
					curBoard.winningBoard = true
				}
			}
		}
	}
}

func part1() {
	calledNumbers, allMatrices := parseFile()
	//fmt.Printf("%#v\n", calledNumbers)
	//fmt.Printf("%#v\n", allMatrices)

	var fastestBoard BingoBoard

	for _, bingoBoard := range allMatrices {
		//fmt.Printf("Checking New Board\n")
		// Calling numbers in order
		for _, num := range calledNumbers {
			if bingoBoard.winningBoard != true {
				// Inc called numbers for the board
				bingoBoard.stepsToWin = bingoBoard.stepsToWin + 1
			}
			for rowPos := 0; rowPos < len(bingoBoard.board); rowPos++ {
				for colPos := 0; colPos < len(bingoBoard.board[rowPos]); colPos++ {
					// If we're not already a winner
					if !bingoBoard.winningBoard {
						if bingoBoard.board[rowPos][colPos].number == num {
							bingoBoard.board[rowPos][colPos].marked = true
						}
						checkWin(&bingoBoard)
						if bingoBoard.winningBoard == true {
							bingoBoard.winningNum = num
						}
					}
				}
			}
		}

		// Check to see if a winning board is the fastest
		if (bingoBoard.winningBoard && bingoBoard.stepsToWin < fastestBoard.stepsToWin) || (bingoBoard.winningBoard && fastestBoard.stepsToWin == 0) {
			//fmt.Printf("    New fastest board found\n")
			fastestBoard = bingoBoard
		}
	}

	for rowPos := 0; rowPos < len(fastestBoard.board); rowPos++ {
		for colPos := 0; colPos < len(fastestBoard.board[rowPos]); colPos++ {
			// If we're not already a winner
			if !fastestBoard.board[rowPos][colPos].marked {
				fastestBoard.unmarkedSum = fastestBoard.unmarkedSum + fastestBoard.board[rowPos][colPos].number
			}
		}
	}

	//fmt.Printf("Winning Board: %#v\n", fastestBoard)
	boardScore := fastestBoard.unmarkedSum * fastestBoard.winningNum
	fmt.Printf("Part 1 Board Score: %#v\n", boardScore)
}

func part2() {
	calledNumbers, allMatrices := parseFile()
	//fmt.Printf("%#v\n", calledNumbers)
	//fmt.Printf("%#v\n", allMatrices)

	var slowestBoard BingoBoard

	for _, bingoBoard := range allMatrices {
		//fmt.Printf("Checking New Board\n")
		// Calling numbers in order
		for _, num := range calledNumbers {
			if bingoBoard.winningBoard != true {
				// Inc called numbers for the board
				bingoBoard.stepsToWin = bingoBoard.stepsToWin + 1
			}
			for rowPos := 0; rowPos < len(bingoBoard.board); rowPos++ {
				for colPos := 0; colPos < len(bingoBoard.board[rowPos]); colPos++ {
					// If we're not already a winner
					if !bingoBoard.winningBoard {
						if bingoBoard.board[rowPos][colPos].number == num {
							bingoBoard.board[rowPos][colPos].marked = true
						}
						checkWin(&bingoBoard)
						if bingoBoard.winningBoard == true {
							bingoBoard.winningNum = num
						}
					}
				}
			}
		}

		// Check to see if a winning board is the fastest
		if (bingoBoard.winningBoard && bingoBoard.stepsToWin > slowestBoard.stepsToWin) || (bingoBoard.winningBoard && slowestBoard.stepsToWin == 0) {
			//fmt.Printf("    New slowest board found\n")
			slowestBoard = bingoBoard
		}
	}

	for rowPos := 0; rowPos < len(slowestBoard.board); rowPos++ {
		for colPos := 0; colPos < len(slowestBoard.board[rowPos]); colPos++ {
			// If we're not already a winner
			if !slowestBoard.board[rowPos][colPos].marked {
				slowestBoard.unmarkedSum = slowestBoard.unmarkedSum + slowestBoard.board[rowPos][colPos].number
			}
		}
	}

	//fmt.Printf("Winning Board: %#v\n", fastestBoard)
	boardScore := slowestBoard.unmarkedSum * slowestBoard.winningNum
	fmt.Printf("Part 2 Board Score: %#v\n", boardScore)
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
