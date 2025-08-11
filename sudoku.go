<<<<<<< HEAD

=======
// package main

// import (
// 	"fmt"
// 	"os"
// )
// func checksNumb (board [][]byte, row,column int, num byte) {
// // /check if same number already exists in the same row

// //we loop across the 9 cells if any given cell contains the same number we want to place num we immmediately return false

// //check row

// //check if number already exists in the same column
// //similar logic but here we loop verically

// // check column
// for i:=0; i < 9; r++{
// 	if board[row][c] == num || board[c][col]== num{
// 		return false
// 	}
// }


// //check the area 3 by 3 sub grid if the same number already exists
// //This is math to find the top-left corner

// // eg if row = 5 row % 3 = 2 so start row  = 5 - 2 = 3
// // eg if col = 7 row % 3 = 1 so start row  = 7 - 2 = 6
// // top left corner of thi subgrid is at (3,6)

// startRow:= (row/3) * 3
// startColumn := (col / 3) * 3
// 	for c:= startRow; c < startRow+3; c++{
// 		for r:=startColumn; r < startColumn+3; r++{
// 			if board[c][r] == num {
// 				return false
// 			}
// 		}
// 	}
// 	return true

// // 	How it works:
// // Row check:
// // Goes through the entire row row to make sure the number isn’t already there.

// // Column check:
// // Goes through the entire column col to make sure the number isn’t already there.

// // 3×3 box check:
// // Finds the top-left corner of the 3×3 square the cell belongs to,
// // then checks that small grid for the same number.

// }
// func solveSudoku(board [][]byte) bool {
// 	for row := 0; row < 9; row++ {
// 		for col := 0; col < 9; col++ {
// 			if board[row][col] == '.' {
// 				for num := byte('1'); num <= '9'; num++ {
// 					if isValid(board, row, col, num) {
// 						board[row][col] = num
// 						if solveSudoku(board, count) {
// 							return true
// 						}
// 						board[row][col] = '.'
// 					}
// 				}
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// func hasUniqueSolution(board [][]byte) bool {
// 	count := 0
// 	copyBoard := make([][]byte, 9)
// 	for i := 0; i < 9; i++ {
// 		copyBoard[i] = make([]byte, 9)
// 		copy(copyBoard[i], board[i])
// 	}
// 	solveSudoku(copyBoard, &count)
// 	return count == 1
// }

// func copyBoard(src [][]byte) [][]byte {
// 	dest := make([][]byte, 9)
// 	for i := 0; i < 9; i++ {
// 		dest[i] = make([]byte, 9)
// 		copy(dest[i], src[i])
// 	}
// 	return dest
// }

// func printBoard(board [][]byte) {
// 	for i := 0; i < 9; i++ {
// 		for j := 0; j < 9; j++ {
// 			fmt.Print(string(board[i][j]))
// 			if j < 8 {
// 				fmt.Print(" ")
// 			}
// 		}
// 		fmt.Println("$")
// 	}
// }

//  func main() {
// 	arguments:= os.Args
// 	fmt.Println(len(arguments))

// 	//arguments must be atleat ten including the program name
// 	if len(arguments) != 10 {
// 		fmt.Println("Error")
// 		return
// 	}

// 		//a 2d array that sets the fixed dimensions of an object
// 		var board [9][9]byte


// 	//validate each row must have atleat 9 characters and they must not be the same
// 	for i:=1; i <=9; i++{
// 		args:=os.Args[i+1]
// 		if len(args) != 9 {
// 			// printing an error if the length of a string is not 9
// 			fmt.Printf("Error")
// 			return 
// 		}
// 		// checking if each argument is either a '.' or a number
// 		// for index,value:= range arguments[i]{
// 		// 	if value != '.' || value < '1' || value < '9' {
// 		// 		fmt.Printf("Error: Row %d has invalid charcters: %c\n", i, value)
// 		// 		return 
// 		// 	}

// 		// 	// a 2d array used for fixing a known dimensions of an array
// 		// 	board[i - 1][index] = byte(value) 
// 		// }
// 				row := make([]byte, 9)
// 		for j := 0; j < 9; j++ {
// 			ch := arg[j]
// 			if ch != '.' && (ch < '1' || ch > '9') {
// 				fmt.Println("Error")
// 				return
// 			}
// 			row[j] = '.'
// 		}
// 		board[i] = row
// 			for i := 0; i < 9; i++ {
// 		arg := os.Args[i+1]
// 		for j := 0; j < 9; j++ {
// 			ch := arg[j]
// 			if ch != '.' {
// 				if !isValid(board, i, j, ch) {
// 					fmt.Println("Error")
// 					return
// 				}
// 				board[i][j] = ch
// 			}
// 		}
// 	}
// 		copyForCount := copyBoard(board)
// 	if countSolutions(copyForCount, 2) != 1 {
// 		fmt.Println("Error")
// 		return
// 	}
// 		if !solve(board) {
// 		fmt.Println("Error")
// 		return
// 	}
// 	printBoard(board)

// 	}

// 	fmt.Println("arguments ok")

// 	fmt.Println(arguments)
// }
>>>>>>> restore-files

package main

import (
	"fmt"
	"os"
)



// 1.) function check if anumber can go in a cell it answers one question: "if i put this number here will it break the sudoku rules"
// Row Check  Loops through all 9 cells in the given row and sees if num is already there.

// Column Check  Loops through all 9 cells in the given column and sees if num is already there.

// 3×3 Subgrid Check  Figures out the top-left corner of the 3×3 box, then loops through all cells in that box to see if num is already there.

// If it passes all checks → return true.
// If it fails any check → return false.

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
	dest := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		dest[i] = make([]byte, 9)
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
