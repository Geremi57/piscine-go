package main

import "fmt"

func fibonaci(n int) int {
	prev := 0
	curr := 1
	for i := 0; i < n; i++ {
		// if i <= 0 {
		// 	return i
		// }
		temp := curr
		curr = prev + curr
		prev = temp
		fmt.Println(curr)
	}
	return curr

}
func fibci(n int) int{
	if n == 0 {
		return 0
	}else if n == 1 {
		return 1
	}	
	return fibci(n - 1) + fibci(n -2)
}
func main() {
	// fmt.Println(fibonaci(8))
	fmt.Println(fibci(9))
}