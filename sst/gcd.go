package main

import "fmt"

func gcd(a,b int) int{
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	n1:=42
	n2:= 12
	
	fmt.Println(gcd(n1, n2))
}