package main

import "fmt"

func itoam2(n int) string {
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
	fmt.Println(itoam2(90))
}