package main

import "fmt"

func itoa(s int) string{
	result := ""
	for s > 0 {
		d := s%10

		result = string(d + '0') + result
		s /= 10
	}
	return result
}


func fromto(a, b int) string{
	fin := ""
	for i := a; i <= b; i++ {
		fin += itoa(i)
		if i != b {
			fin += ", "
		} 
	}
	return fin

}

func main() {
	n1 := 1
	n2 := 111
	fmt.Println(fromto(n1, n2))
}