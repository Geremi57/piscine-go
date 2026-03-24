package main

import (
	"fmt"
)


func printHex(b byte) {
	str := "0123456789abcdef"
	min:= b/16
	max:=b%16

	fmt.Print(string(str[min]))
	fmt.Print(string(str[max]))
}

func PrintMemory(arr [10]byte) {
	for i,v := range arr{
		printHex(v)
			fmt.Print(" ")
		if (i+1)%4==0{
			fmt.Print("\n")
		}
	}

		fmt.Println()
	for _,k := range arr{
		fmt.Print(string(k))
		if k < 32 {
		fmt.Print(string('.'))
		}
	}
}

func main(){
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})

}