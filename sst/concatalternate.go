package main

import "fmt"

func ConcatAlternate(slice1, slice2 []int) []int{
	fin := [] int {}
	minlength := len(slice1)
	if len(slice1) >= len(slice2) {
		minlength = len(slice2)
	}
		for i:=0; i < minlength; i++{
			fin = append(fin, slice1[i], slice2[i])

		}
		if len(slice1) > minlength {
			fin = append(fin, slice1[minlength:]...)
		}
		if len(slice2) > minlength {
			fin = append(fin, slice2[minlength:]...)
		}
		return fin
	
}

func main() {
	fmt.Println(ConcatAlternate([]int{2,4,6,8,10}, []int{3,5,7,9,11}))
}