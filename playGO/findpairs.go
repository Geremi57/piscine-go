package main

import (
	"fmt"
)

func findPairs(arr []int, target int) [][]int {
	var result [][]int

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 6

	pairs := findPairs(arr, target)

	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
		return
	}

	fmt.Printf("Pairs with sum %d: %v\n", target, pairs)
}