package main

import "fmt"

func convertToBase(n int, symbols string) string {

	if n == 0 {
	return string(symbols[0])
}

	base := len(symbols)
	result := ""

	isNegative := false

	if n < 0 {
		isNegative = true
		n = -n
	}

	for n > 0 {
		remainder := n % base
		result = string(symbols[remainder])+ result
		n = n / base
	}

	if isNegative{
		result = "-" + result
	}
	return result

}

func main() {
	fmt.Println(convertToBase(125, "0123456789"))
	fmt.Println(convertToBase(125, "aa"))
	fmt.Println(convertToBase(-125, "choumi"))
	fmt.Println(convertToBase(125, "0123456789ABCDEF"))
}