package main

import "fmt"

func fibonacci(n int) int {
	var sum int = 0
	var count int = 1
	for i := 2; i <= n; i++ {
		next := sum + count
		sum = count
		count = next
		// count = i + 1 - 1
		fmt.Println(count)
	}
	return count
}
func main() {
	fibonacci(9)
}

<<<<<<< HEAD
// https://github.com/Frenchris/piscine-go-2/blob/master/raid1a.go
=======
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}
>>>>>>> 4521c38 (sudoku algorithm)
