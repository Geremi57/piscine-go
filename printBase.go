package main

import "fmt"

func convertToBase(n int, symbols string) string {
if n == 0 {
	return string(symbols[0])
}
	base := len(symbols)
	result := ""
	for n > 0 {
		remainder := n % base
		result = string('0' + remainder) + result
		n = n / base
	}
	return result

}

func main() {
	fmt.Println(convertToBase(125, "0109oui"))
}
