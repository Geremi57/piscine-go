package main

import "fmt"
func Itoa(n int) string{
	if n == 0 {
		return "0"
	}

	if n < 0 {
		n = -n
	}

	// /store the digits in reverse
	// digits := []rune{}
	results := ""
	for n > 0{
		//getting the last digit
		d:= n%10
		results = string(d + '0') + results
		n /= 10
	}
	return results
}

func main() {
	fin:= 90
	ans:= Itoa(fin)
	fmt.Printf("The type of %v is a %T", fin, ans)

}

