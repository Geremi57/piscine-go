package main

import "fmt"

func itoa(n int) string{
	if n == 0 {
		return "0"
	}
	if n < 0 {
		n = -n
	}
	res := ""

	for n > 0 {
		d := n%10

		res = string(d + '0') + res
		n /= 10
	}

	return res
}

func main() {
	fmt.Println(itoa(13))
}