package main

import "fmt"

func concatalt(sl1 []int, sl2 []int) []int{
	fin := []int{}
	minlength := len(sl1)

	if len(sl1) > len(sl2) {
		minlength = len(sl2)
	}

	for i:=0; i < minlength; i++{
		fin=append(fin, sl1[i], sl2[i])
	}
	if len(sl1) > minlength {
		fin = append(fin, sl1[minlength:]...)
	}
	if len(sl2) > minlength {
		fin = append(fin, sl2[minlength:]...)
	}

	return fin
}

func revConAlt(sl1 []int, sl2 []int) []int{
	final := []int{}
	minlen := len(sl1)
	if len(sl1) > len(sl2) {
		minlen = len(sl2)
	}
	for i:=0; i < minlen; i++{
		final = append(final, sl1[i], sl2[i])
	}

	if len(sl2) > minlen {
		final = append(final, sl2[minlen:]...)
	}
	if len(sl1) > minlen {
		final = append(final, sl1[minlen:]...)
	}
	result := []int{}

	for i:=len(final) - 1; i >= 0; i -- {
		result = append(result, final[i])
	}
	return result
}
func main() {
	fmt.Println(revConAlt([]int{2,4,6,8}, []int{3,5,7,9,10}))
		fmt.Println(concatalt([]int{2,4,6,8}, []int{3,5,7,9,10}))
		
}