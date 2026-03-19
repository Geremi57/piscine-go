package main

import "fmt"

func atoi(str string) int{
	n := 0
	for _,s := range str {
		if s < '0' && s > '9' {
			return 0
		}
		n = n * 10 + int(s - '0')
	}
return n
}

for n > 0{
		//getting the last digit
		// d:= n%10
		results = string(n%10 + '0') + results
		n /= 10
	}

func main() {
	// fmt.Println("9" + 9)
	fmt.Println(atoi("9") + 9)
}