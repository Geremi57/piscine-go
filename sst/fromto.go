<<<<<<< HEAD
package main

import "fmt"

func itoa(s int) string{
	result := ""
	for s > 0 {
		d := s%10

		result = string(rune(d) + '0') + result
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
	n1 := 12
	n2 := 56
	fmt.Println(fromto(n1, n2))
=======
package main

import "fmt"

func Itoa(n int) string{
	results := ""
	if n == 0 {
		results = "0"
	}

	for n > 0{
		d:=n%10
		results = string(rune(d) + '0') + results
		n /= 10
	}
	return results
}

func FromTo(from int, to int) string {
	final:= ""
	if from < 0 || to < 0 || from > 99 || to  > 99{
		final = "Invalid\n"
	}
	for i:=from; i < to; i++{
		final += Itoa(i) + final
	}
	
	return final
}

func main() {
	fmt.Println(FromTo(12, 56))
>>>>>>> 5541715 (write md for first question)
}