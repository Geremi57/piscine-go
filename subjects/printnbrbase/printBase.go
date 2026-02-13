package main

import "fmt"

func printBase(n int, symbols string) string{
	if n == 0 {
		return string(symbols[0])
	}
	result := ""
	base := len(symbols)
	isNegative := false
	
	if n < 0 {
		isNegative = true
		n = -n
	}

	for n > 0 {
		rem := n % base
		result  = string(symbols[rem]) + result
		n = n / base
	}

	if isNegative {
		result = "-" + result
	}

	return result
}

func main() {
	fmt.Println(printBase(125, "0123456789"))
	fmt.Println(printBase(125, "aa"))
	fmt.Println(printBase(-125, "choumi"))


}