package main

import "fmt"

func RevConcatAlternate(slice1,slice2 []int) []int {
	var res []int
	i:=len(slice1) -1
	j:=len(slice2) -1

	if len(slice1) > len(slice2){
		for i >= len(slice2){
			res = append(res, slice1[i])
			i--
		}
	}else if len(slice2) > len(slice1){
		for j >= len(slice1) {
			res = append(res, slice2[j])
			j--
		}
	}

	for i >=0 && j >= 0 {
		res = append(res, slice1[i])
		i--

		res = append(res, slice2[j])
		j--
	}

return res
}

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}