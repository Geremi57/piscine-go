package main

import (
	"fmt"
)

func Itoa(n int) string{
	res := ""
	for n > 0 {
		d := n%10
		res = string(d + '0') + res
		n /= 10
	}
	return res
}

func ZipString(s string) string {
	count := 1
	result := ""
	for i:=1; i < len(s); i++{
		if s[i] == s[i-1] {
			count++
		}else{
			count = 1
			result += Itoa(count)+string(s[i-1])
		}
	}
	result += Itoa(count) + string(s[len(s) -1 ])
	return result
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}