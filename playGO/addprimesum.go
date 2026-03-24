package main

import (
	"os"
	"fmt"
)

func atoi(s string) int{
	n := 0
	for _,v := range s{
		if 	v < '0' || v > '9' {
			return 0
		}
		n = int(v - '0') + n * 10
	}
	return n
}

func isPrime(n int)bool{
	for i :=2; i*i <= n; i++{
		if n % i == 0{
			return false
		}
	}
	return true
}

func addPrimeSum(n int) int{
	count := 0
	for i:=2; i <= n; i++{
		if isPrime(i) {
			count+=i
		}
	}
	return count
}

func main(){
	args := os.Args[1]
	numb := atoi(args)
	
	fmt.Println(addPrimeSum(numb))
}