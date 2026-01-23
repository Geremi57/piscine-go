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

func main() {
	// fmt.Println("9" + 9)
	fmt.Println(atoi("9") + 9)
}