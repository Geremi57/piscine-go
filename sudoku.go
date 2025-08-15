package main

import (
	"fmt"
	"os"
)

func isValid(board [][]byte, row, col int, num byte) bool {
	// check row and column
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}
	// check 3x3 subgrid
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			if board[r][c] == num {
				return false
			}
		}
	}
	return true
}

// 2.) REcursive backtracking(the brain)
// This is the brain of the Sudoku solver. It works like:

// Find an empty cell ('.').

// Try placing numbers '1' through '9' there.

// For each number:

// Call isValid to check if it’s allowed.

// If valid, place it and recursively call solve to try to solve the rest.

// If recursion succeeds → puzzle solved → return true.

// If recursion fails → reset the cell to '.' (backtrack) and try the next number.

// // If no numbers work → return false (dead end).


// think of it as 
// function solve(board) {
//   if (board is solved) return true;
//   else {
//     try a number;
//     if it works, keep going;
//     if not, erase it and try next number;
//   }
// // }

func solve(board [][]byte) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				for num := byte('1'); num <= '9'; num++ {
					if isValid(board, r, c, num) {
						board[r][c] = num
						if solve(board) {
							return true
						}
						board[r][c] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

// Works like solve, but instead of stopping when one solution is found, it counts how many solutions there are.

// If count ever reaches limit (in this case 2), it stops early — no need to keep searching.

// Used to make sure the given puzzle has exactly 1 solution.

// This is a backtracking solver, but instead of stopping early, it counts all solutions until limit is reached.

// That return 1 at the bottom is the base case it means "This path is a complete, valid solution".

func countSolutions(board [][]byte, limit int) int {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				count := 0
				for num := byte('1'); num <= '9'; num++ {
					if isValid(board, r, c, num) {
						board[r][c] = num
						count += countSolutions(board, limit)
						board[r][c] = '.'
						if count >= limit {
							return count
						}
					}
				}
				return count
			}
		}
	}
	return 1
}

// Creates a new 9×9 array and copies all values from src.

// Necessary because countSolutions modifies the board while checking.


func copyBoard(src [][]byte) [][]byte {
    // Step 1: Create a new outer slice (like `let dest = []`)
    dest := make([][]byte, 9)

    // Step 2: For each row
    for i := 0; i < 9; i++ {
        // create a new row of 9 bytes (like `new Array(9)`)
        dest[i] = make([]byte, 9)

        // copy the values from src[i] into dest[i]
        copy(dest[i], src[i])
    }

    return dest
}


// Loops through each row and column, printing values with spaces.

// Outputs the final solved Sudoku in a nice format.

func printBoard(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c", board[i][j])
			if j < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// main


// Argument Check

// Needs exactly 10 arguments:

// Arg 0 → program name

// Args 1–9 → 9 strings representing Sudoku rows.

// Format Check

// Each row must be exactly 9 characters.

// Each character must be '.' or '1'–'9'.

// If invalid → print "Error" and exit.

// Load Initial Board

// Creates a 9×9 board filled with '.'.

// Initial Validity Check

// Places given numbers one by one.

// Calls isValid before placing each — if a number already conflicts, print "Error".

// Unique Solution Check

// Calls countSolutions(copyBoard(board), 2).

// If the result isn’t 1, print "Error".

// Solve Puzzle

// Calls solve(board).

// If unsolvable → "Error".

// Print Result

// Calls printBoard to display the solved puzzle.



func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	board := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		arg := os.Args[i+1]
		if len(arg) != 9 {
			fmt.Println("Error")
			return
		}
		row := make([]byte, 9)
		for j := 0; j < 9; j++ {
			ch := arg[j]
			if ch != '.' && (ch < '1' || ch > '9') {
				fmt.Println("Error")
				return
			}
			row[j] = '.'
		}
		board[i] = row
	}

	for i := 0; i < 9; i++ {
		arg := os.Args[i+1]
		for j := 0; j < 9; j++ {
			ch := arg[j]
			if ch != '.' {
				if !isValid(board, i, j, ch) {
					fmt.Println("Error")
					return
				}
				board[i][j] = ch
			}
		}
	}

	if countSolutions(copyBoard(board), 2) != 1 {
		fmt.Println("Error")
		return
	}

	if !solve(board) {
		fmt.Println("Error")
		return
	}

	printBoard(board)
}
