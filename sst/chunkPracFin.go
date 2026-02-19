package main

import "fmt"


func Chunk(slice []int, size int) {
	var main [][]int

	if size <= 0 {
		fmt.Println()
		return
	}
	
	for i:=0; i < len(slice); i+=size{
		end := i+size

		if end > len(slice){
			end = len(slice)
		}

		main = append(main, slice[i:end])
	}

	fmt.Println(main)
}

func main() {

Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}