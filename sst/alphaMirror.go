package main

import (
	"fmt"
	"os"
)

func alphaMirror(s string) string {
	final := ""
	uStart := []rune{}
	uFinal := []rune{}
	lStart := []rune{}
	lFinal := []rune{}
	
	for j := 'Z'; j >= 'A'; j-- {
		uStart = append(uStart, rune(j))
	}
	for i:= 'A'; i <= 'Z'; i++{
		uFinal = append(uFinal, rune(i))
	}
	for j := 'z'; j >= 'a'; j-- {
		lStart = append(lStart, j)
	}
	for i:= 'a'; i <= 'z'; i++{
		lFinal = append(lFinal, i)
	}

for _,value := range s{
	if value >= 'a' && value <= 'z'{
		for j,c := range lStart{
				if value == c {
					final += string(lFinal[j])
				}
		}
	}else if value >= 'A' && value <= 'Z'{
			for i, v := range uStart {
					if value == v{
						final += string(uFinal[i])
					}
			}
	}else{
		final += string(value)
	}
}
	return final
}

func main() {
	arg := os.Args[1:]
	str := ""
	for _,v := range arg{
		str += string(v)
		str += string(' ')
	}
	fmt.Println(alphaMirror(str))
}
