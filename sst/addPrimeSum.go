package main

import ("os"
"fmt")

func Atoi(s string) int{
	fin:=0
	for _, value := range s {
		if value < '0' || value > '9'{
			return 0
		}
		fin = fin * 10 + int(value - '0')
	}
	return fin
}

func isPrime(n int) bool{
	if n < 2 {
		return false
	}
	for i:=2; i*i <= n; i++{
		if n % i == 0 {
			return false
		}
	}
	return true

}

func addPrimeSum(n int) int{
	fin := 0
	for i:=2; i <= n; i++{
			if isPrime(i){
			fin += i 
		}
	}
	return fin

}

func main() {
	if len(os.Args) > 2 {
    fmt.Println("0")
	return
}
	arg := os.Args[1]
	fmt.Println(addPrimeSum(Atoi(arg)))
}