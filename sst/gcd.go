package main

import (
	"fmt"
	"os"
)

func gcd(a,b int) int{
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func atoi(str string) int{
	n:=0
	for _,v := range str {
		if v < '0' || v > '9' {
			return 0
		}
		n = n * 10 + int(v - '0')
	}
	return n
}

func itoa(i int) string{
	if i == 0 {
		return "0"
	}
	if i < 0 {
		i = -i
	}
res := ""
	for i > 0 {
		d := i%10
		res = string(d + '0') + res
		i /= 10		
	}
	return res
}

func main() {
	// name := os.Args[0]
	n1 := os.Args[1]
	n2:=os.Args[2]
	// n2:= os.Args[3]
	fmt.Println(n1, n2)
	fmt.Printf("this is a %t", n1)
	fmt.Println(itoa(atoi(n1)), itoa(atoi(n2)))
	fmt.Println(gcd(atoi(n1), atoi(n2)))
}