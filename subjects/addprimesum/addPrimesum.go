package main

import (
	"fmt"
	"os"
)

func atoi(str string) int{
	n := 0 
	
	for _,v := range str{
		if v < '0' || v > '9' {
			return 0
		}
		n = n * 10 + int(v - '0')
	}
	return n
}

func isPrime(n int) bool{
	 if n < 2 {
		return false
	 }

	 for i:=2; i*i <= n; i++{
		if n%i == 0 {
			return false
		}
	 }
	 return true
}

func addPrimeSum(n int) int{
	
	fin := 0
	for i:=2; i <= n; i++ {
		if isPrime(i) {
			fin += i
		}
	}
	return fin
}

func main() {
	args := os.Args[1]

	if len(args) < 1 {
		fmt.Println(" ")
	}
	// args := os.Args[1]
	fmt.Println(addPrimeSum(atoi(args)))
}