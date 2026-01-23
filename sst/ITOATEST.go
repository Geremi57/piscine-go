package main

import "fmt"

func itoa(i int) string{
	// n := 0
	if i == 0 {
		return "0"
	}
	if i < 0 {
		i = -i
	}
	results := ""
	for i > 0 {
		d := i%10
		results = string(d + '0') + results
		i /= 10 

	}
	return results
}

func atoi(str string) int{
	n := 0
	for _,v := range str{

		if v < '0' ||  v > '9' {
			return 0
		}
		n = n * 10 + int(v - '0')
	}
	return n
}

func main() {
	fmt.Println(itoa(123))
	fmt.Println(atoi("young"))
}