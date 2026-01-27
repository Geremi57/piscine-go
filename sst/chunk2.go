package main

import "fmt"

func chunk(slice []int, n int){

	if n <= 0 {
		fmt.Println()
		return
	}

	var mainS [][]int

	for i:=0; i < len(slice); i+=n {

		end := i + n
		if end > len(slice){
			end = len(slice)
		}

		chunk := slice[i:end]

		mainS = append(mainS, chunk)
	}
	fmt.Println(mainS)
}

func main() {
		sl := [] int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		chunk(sl, 8)

}