package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func isValid(board [][]byte, row, col int, n byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == n || board[i][col] == n {
			return false
		}
	}
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == n {
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

func solveHandler(w http.ResponseWriter, r *http.Request) {
	var board [][]string
	if err := json.NewDecoder(r.Body).Decode(&board); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	b := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		b[i] = make([]byte, 9)
		for j := 0; j < 9; j++ {
			if board[i][j] == "" {
				b[i][j] = '.'
			} else {
				b[i][j] = board[i][j][0]
			}
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			ch := b[i][j]
			if ch != '.' {
				b[i][j] = '.' // clear it
				if !isValid(b, i, j, ch) {
					http.Error(w, "Invalid board", http.StatusBadRequest)
					return
				}
				b[i][j] = ch // restore
			}
		}
	}

	if !solve(b) {
		http.Error(w, "Unsolvable puzzle", http.StatusBadRequest)
		return
	}

	res := make([][]string, 9)
	for i := 0; i < 9; i++ {
		res[i] = make([]string, 9)
		for j := 0; j < 9; j++ {
			res[i][j] = string(b[i][j])
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	// Serve static files (index.html, style.css, script.js)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// API endpoint
	http.HandleFunc("/solve", solveHandler)

	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	fmt.Println("Server running on http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
