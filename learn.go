package main

import (
	"fmt"
)

func main() {
	n:=10
	j:=3
	// var n bool = true
	fmt.Println(n & j )
	fmt.Println(n | j )
	fmt.Println(n ^ j )
	fmt.Println(n &^ j )
	fmt.Printf("%v %T" ,n, n)
}