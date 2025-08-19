package main

import "fmt"

func Atoi(s string) int{
	n := 0
	for _,v := range s {
if v < '0' || v > '9' {
	return 0
}

n = n * 10 + int(v - '0')
	}
	return n
}

func main() {
	s:= "9"
	fin:= Atoi(s)
	fmt.Printf("this is an %T\n", s)
	fmt.Printf("this is an %T\n", fin)
}

