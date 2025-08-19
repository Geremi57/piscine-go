package main

import "fmt"

func concatSlice(slice1, slice2 []int) []int{
	fin := []int{}
	for i:=0; i < len(slice1); i++{
		fin = append(fin, slice1[i])
	}
	for i:=0; i < len(slice2); i++{
		fin = append(fin, slice2[i])
	}

	return fin
}

func main() {
	fmt.Println(concatSlice([]int{1,2,3}, []int{}))
}