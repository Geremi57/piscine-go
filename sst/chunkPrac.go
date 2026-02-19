package main

import "fmt"

func chunk(slice []int, size int){
	// var/
	var main [][]int

for i := 0 ; i < len(slice); i+=size{
	end := i + size
	if end > len(slice){
		end = len(slice)
	}

	chunk := slice[i:end]
	main = append(main, chunk) 
}

fmt.Println(main)

}

func main(){
	chunk([]int{1, 2, 3, 4, 5, 7, 8, 9}, 6)
	chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}