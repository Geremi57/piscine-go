package main

import "fmt"

func FifthAndSkip(str string) string{
	final:= ""

	for _, value:= range str{
		if value  >= 'A' && value <= 'Z' || value >= 'a' && value <= 'z' || value >= '0' && value  <= '9'{
				final += string(value)
		}
	}
	s:=""
	for i:=0; i < len(final); i++{
		s += string(final[i])
			if (i+1) % 5 == 0{
				s += " "
			}
	}
	return s

}

func main() {
	fmt.Println(FifthAndSkip("abcd efgh ijkl mnop qrst"))
}