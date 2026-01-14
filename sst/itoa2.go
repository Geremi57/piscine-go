package main

import "fmt"

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	// First start by innotializing an empty string where our solution will be 
	results := ""
	for n > 0 {
		// check for numbers greater than 0 we need to find the last digit first eg 5 = 125
		d := n%10
		// turn the number into a rune first then add by the '0' which will give the number a rune value 
		// then finally converting to a string and prepending it by adding the results
		results = string(rune(d) + '0') + results
		//meaning number is number / 10 reducing the size of n after every iteration
		n/=10
	}
	fmt.Println(results)
	return results
}

func main() {
	num:= 99
	fmt.Printf("%v is a %T\n", num, num)
	val := Itoa(num)
	fmt.Printf("%v is a %T\n", val, val)

}