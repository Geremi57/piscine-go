package main

import ("os"
 "fmt")

func isValid(board [][]byte, row, col int, n byte) bool{
	for i := 0; i < 9; i++ {
		if board[row][i] == n || board[i][col] == n{
			return false
		}
	} 
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow + 3; i++{
		for j := startCol; j < startCol + 3; j++ {
			if board[i][j] == n {
				return false
			}
		}
	}
	return true
}

func solve(b [][]byte) bool{
	for r := 0; r < 9; r++{
		for c := 0; c < 9; c++{
			if b[r][c] == '.'{
				for n := byte('1'); n <= '9'; n++{
					if isValid(b, r,c, n){
						b[r][c] = n
						if solve(b){
							return true
						}
						b[r][c] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func countSolutions(board [][]byte, limit int) int{
	for r := 0; r < 9; r++{
		for c:= 0; c < 9; c++{
			if(board[r][c] == '.'){
				count := 0
				for num := byte('1'); num <= '9'; num++{
					
					if (isValid(board, r, c, num)){
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

func copyBoard(src [][]byte) [][]byte{
	dest := make([][]byte, 9)
for i := 0; i < 9; i++{
	dest[i]= make([]byte, 9)
	copy(dest[i], src[i])
}
return dest
}

func printBoard(board [][]byte){
	for i := 0; i < 9; i++{
		for j := 0; j < 9; j++{
			fmt.Printf("%c", board[i][j])
			if j < 8{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main(){
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
	for j := 0; j < 9; j ++{
		ch := arg[j]
		if ch != '.' && (ch < '1' || ch > '9'){
			fmt.Println("Error")
			return
		}
		row[j] = '.'
	}
	board[i] = row
	}
	for i:=0; i < 9; i++{
		arg := os.Args[i+1]
		for j := 0; j < 9; j++{
			ch := arg[j]
			if ch != '.'{
				if !isValid(board, i , j, ch){
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

	if !solve(board){
		fmt.Println("Error")
		return
	}

	printBoard(board)
}
