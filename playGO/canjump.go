package main

import "fmt"

func CanJump(arr []uint)bool{
	i := 0
	for i < len(arr){
		if i == len(arr)-1{
			return true
		}
		step := int(arr[i])
		if step == 0 {
			return false
		}
		i += step
		if i >= len(arr){
			return false
		}
	}
	return false
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}