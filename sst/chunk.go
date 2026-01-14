<<<<<<< HEAD
package main

import "fmt"

func Chunk(slice []int, size int) {
	if size <= 0 {
		fmt.Println()
		return 
	}
	var mainS [][]int

for i:=0; i < len(slice); i += size{

	end := i + size
	if end > len(slice) {
		end = len(slice)
	}

	Chunk:= slice[i:end]

	mainS = append(mainS, Chunk)
}
fmt.Println(mainS)

}

func main() {
	sl := [] int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	Chunk(sl, 0)
=======
package main

import "fmt"

func Chunk(slice []int, size int) {
if size == 0 {
	return fmt.Println()
}


}

func main() {

>>>>>>> 5541715 (write md for first question)
}