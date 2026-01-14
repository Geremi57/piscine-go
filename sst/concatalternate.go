<<<<<<< HEAD
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
=======
package main

import "fmt"

func ConcatAlternate(slice1, slice2 []int) []int {
	slice3 := []int{}
	l := len(slice1) + len(slice2)
	// long := len(slice1) > lien(slice2)
	if len(slice1) >= len(slice2) {
		// slice3[0] = slice1[0]
		slice3 = append(slice3, slice1[0])

	} else if len(slice1) < len(slice2) {
		slice3 = append(slice3, slice2[0])
	}
	for i, k := 0, 1; len(slice3) < l ; {
		if i < len(slice1) {
			slice3 = append(slice3, slice1[i])
			i++
		}
		if k < len(slice2){
			slice3 = append(slice3, slice2[k])
			k++
		}
	}
	return slice3
}

func main() {
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
>>>>>>> 5541715 (write md for first question)
}