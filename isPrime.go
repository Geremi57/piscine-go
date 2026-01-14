hrmupackage main

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false 
	}
	for i := 2; i*i <= n; i++ {
		// if n != i && i != 1{
			if n%i == 0 {
				return false
			}

	}
	return true
}

func main() {
	n := 31
	fmt.Print(isPrime(n))
}
