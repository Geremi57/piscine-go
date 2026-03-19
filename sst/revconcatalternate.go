package main

import "fmt"


func RevConcatAlternate(slice1, slice2 []int) []int {
	result := []int{}

	i := len(slice1) - 1
	j := len(slice2) - 1

	// Phase 1: handle extra elements
	for i > j {
		result = append(result, slice1[i])
		i--
	}
	for j > i {
		result = append(result, slice2[j])
		j--
	}

	// Phase 2: alternate (equal sizes now)
	for i >= 0 && j >= 0 {
		result = append(result, slice1[i])
		result = append(result, slice2[j])
		i--
		j--
	}

	return result
}
func main() {
	fmt.Println(RevConcatAlternate([]int{1,2,3}, []int{4,5,6,7,8}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}