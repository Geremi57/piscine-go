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

// https://github.com/Frenchris/piscine-go-2/blob/master/raid1a.go
