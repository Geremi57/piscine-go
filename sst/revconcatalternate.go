package main

import "fmt"

func RevConcatAlternate(slice1, slice2 []int) []int {
	fin := []int{}
	res := []int{}
	minlength:= len(slice1)

	if minlength > len(slice2) {
		minlength = len(slice2)
	}

	for i := 0; i < minlength; i++ {
		fin = append(fin, slice1[i], slice2[i])
	}
	// return fin
	if minlength < len(slice1) {
		fin = append(fin, slice1[minlength:]...)
	}
	if minlength < len(slice2) {
		fin = append(fin, slice2[minlength:]...)
	}
	
	for i:=len(fin) - 1; i >= 0; i--{
		res = append(res, fin[i])
	}
	return res

}

func main() {
	fmt.Println(RevConcatAlternate([]int{1,2,3}, []int{4,5,6,7,8}))
}