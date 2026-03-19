package main

import "fmt"

func Chunk(slice []int, size int) {
	if size <= 0 {
		fmt.Println()
		return
	}

	var mains [][]int

	for i:=0; i < len(slice); i+=size{
		end := i + size
		if end > len(slice){
			end = len(slice)
		}
		mains = append(mains, slice[i:end])
	}
	fmt.Println(mains)

	// return mains
}

func main() {
sl := [] int{0, 1, 2, 3, 4, 5, 6, 7}
	Chunk(sl, 0)
Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
