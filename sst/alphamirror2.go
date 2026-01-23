package main

import (
	"fmt"
	"os"
)

func mirror(s string) string {
	// arguments := os.Args[1:]
	// args := arguments[1]
	final := ""
	uStart := []rune{}
	uFinal := []rune{}
	lStart := []rune{}
	lFinal := []rune{}


	
	for j := 'Z'; j >= 'A'; j-- {
		uStart = append(uStart, j)
	}
	for i:= 'A'; i <= 'Z'; i++{
		uFinal = append(uFinal, i)
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
	fmt.Println(final)
	return final

}

func main() {
	arguments := os.Args[1:]
	str := ""
	for _,c := range arguments {
		str += string(c)
		str += string(' ')
	}
	fmt.Println(mirror(str))

	// mirror("my head")
}