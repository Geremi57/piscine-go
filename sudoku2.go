package main

import (
	"fmt"
	"os"
)

func isValid(b [][]byte, r, c int, n byte) bool {
	for i := 0; i < 9; i++ {
		if b[r][i] == n || b[i][c] == n {
			return false
		}
	}
	sr, sc := (r/3)*3, (c/3)*3
	for i := sr; i < sr+3; i++ {
		for j := sc; j < sc+3; j++ {
			if b[i][j] == n {
				return false
			}
		}
	}
	return true
}

func solve(b [][]byte) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if b[r][c] == '.' {
				for n := byte('1'); n <= '9'; n++ {
					if isValid(b, r, c, n) {
						b[r][c] = n
						if solve(b) {
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

func count(b [][]byte, lim int) int {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if b[r][c] == '.' {
				cnt := 0
				for n := byte('1'); n <= '9'; n++ {
					if isValid(b, r, c, n) {
						b[r][c] = n
						cnt += count(b, lim)
						b[r][c] = '.'
						if cnt >= lim {
							return cnt
						}
					}
				}
				return cnt
			}
		}
	}
	return 1
}

func copyB(s [][]byte) [][]byte {
	d := make([][]byte, 9)
	for i := range s {
		d[i] = append([]byte{}, s[i]...)
	}
	return d
}

func printB(b [][]byte) {
	for i := range b {
		for j := range b[i] {
			fmt.Printf("%c", b[i][j])
			if j < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}
	b := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		s := os.Args[i+1]
		if len(s) != 9 {
			fmt.Println("Error")
			return
		}
		b[i] = []byte(".........")
		for _, ch := range s {
			if ch != '.' && (ch < '1' || ch > '9') {
				fmt.Println("Error")
				return
			}
		}
	}
	for i := 0; i < 9; i++ {
		for j, ch := range os.Args[i+1] {
			if ch != '.' {
				if !isValid(b, i, j, byte(ch)) {
					fmt.Println("Error")
					return
				}
				b[i][j] = byte(ch)
			}
		}
	}
	if count(copyB(b), 2) != 1 || !solve(b) {
		fmt.Println("Error")
		return
	}
	printB(b)
}
