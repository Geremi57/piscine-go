package main

import "fmt"

func HashCode(dec string) string{
	l := len(dec)
	res := ""
	for _,c := range dec {
		v := (int(c) + l)%127
	if v < 33 {
		v += 33
	}
	res += string(rune(v))
	}
	return res

}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AC"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
	fmt.Println(HashCode("14 avril 2014"))

}