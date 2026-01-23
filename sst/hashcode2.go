package main

import "fmt"

func HashCode(str string) string {
	res := ""
	l := len(str)
	for _,c := range str {
		v := (int(c) + l)%127
		
		if v < 33 {
			v += 33
		}
		res += string(v)
	}
	return res

}
func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
	fmt.Println(HashCode("14 Avril 2014"))

}