package main

import (
	"os"

	"github.com/01-edu/z01"

	// piscine ".."
)

var table [9][9]int

func main() {

	lengthArgs := piscine.ArrayStrLength(os.Args)
	if lengthArgs != 10 {
		piscine.PrintStr("Error")
		z01.PrintRune('\n')
	} else {
		table = parseInput()
		if backtrack(&table) {
			printBoard(table)
		} else {
			piscine.PrintStr("Error")
			z01.PrintRune('\n')
		}
	}
}

//Prints the sudoku board
func printBoard(board [9][9]int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if col == 8 {
				piscine.PrintNbr(table[row][col])
			} else {
				piscine.PrintNbr(table[row][col])
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

//transforms the input to board form
func parseInput() [9][9]int {
	board := [9][9]int{}

	row := 0
	col := 0

	for i := 1; i <= 9; i++ {
		chars := []rune(os.Args[i])
		for j := 0; j < piscine.StrLen(string(chars)); j++ {
			if chars[j] == '.' {
				board[row][col] = 0
			}
			board[row][col] = piscine.Atoi(string(chars[j]))
			col++
			if col == 9 {
				col = 0
			}
		}
		row++
	}

	return board
}

//checks if the board is valid
func isBoardValid(board *[9][9]int) bool {

	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

//checks if an array has duplicates
func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func backtrack(board *[9][9]int) bool {

	if !(hasEmptyCell(board)) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					board[i][j] = candidate
					if isBoardValid(board) {
						if backtrack(board) {
							return true
						}
						board[i][j] = 0
					} else {
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func hasEmptyCell(board *[9][9]int) bool {

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}