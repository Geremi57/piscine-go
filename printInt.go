package main

import (
	"fmt"
)

func main() {

 var arr = [3]int{0,0,0}
 
	for i:=0; i < 7; i++{
		if arr[0] <= i+2{
			arr[1] = i+2
		}
		if arr[1] <= i+3{
			arr[2] = i+3
		}

		arr[0]++
		result := 0
		result = arr[0] *10 + arr[1] 
		res:= result * 10+ arr[2]
		fmt.Print(res,",")
	}
}
