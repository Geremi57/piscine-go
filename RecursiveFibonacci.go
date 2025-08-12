package main

import "fmt"

func fibonacci(n int) int {
	var sum int = 0
	var count int = 1
		if n < 0 {
			return - 1
		}
		if n == 1 {
			return 0
		}
		if n == 1 {
			return 1
		}

		return fibonacci(n - 1) + fibonacci(n - 2)
	
	// return count
}
func main() {
	fmt.Println(fibonacci(9))
}
