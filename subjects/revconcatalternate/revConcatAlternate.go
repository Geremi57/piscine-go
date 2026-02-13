package main

import "fmt"

func concatAlternate(sl1 []int, sl2 []int)[]int {
	minlength := len(sl1)
	fin := []int{}
	reversed := []int{}

	if minlength > len(sl2){
		minlength = len(sl2)
	}

	for i:=0; i < minlength; i++{
		fin = append(fin, sl1[i], sl2[i])
	}

	if len(sl1) > minlength{
		fin = append(fin, sl1[minlength:]...)
	} 
	if len(sl2) > minlength{
		fin = append(fin, sl2[minlength:]...)
	} 

	for i:= len(fin) -1; i >= 0; i--{
		reversed = append(reversed, fin[i])

	}
	return reversed
}

func main(){
	sl := []int{1,3,5,7,9}
	sl2 := []int{2,4,6,8,10}
	fmt.Println(concatAlternate(sl, sl2))
}
