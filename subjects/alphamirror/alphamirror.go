package main

import (
	"fmt"
	"os"
)

func alphaMirror(str string) string{
	uStart:=[]rune{}
	uReverse:=[]rune{}
	lStart := []rune{}
	lReverse := []rune{}
	fin := ""


	for i:='Z'; i >= 'A'; i--{
		uReverse = append(uReverse, i)
	}
	for j:='A'; j <= 'Z'; j++{
		uStart = append(uStart, j)
	}

	for k:='a'; k <= 'z'; k++{
		lStart = append(lStart, k)
	}

	for m:='z'; m >= 'a'; m--{
		lReverse = append(lReverse, m)
	}
	
	for _,v := range str {
		if v >= 'A' && v <= 'Z' {
			for i,j := range uStart{
				if v == j{
					fin += string(uReverse[i])
				}
			}
		}else if v >= 'a' && v <= 'z' {
			for i,j := range lStart{
				if v == j {
					fin += string(lReverse[i])
				}
			}
		}else {
			fin += string(v)
		}
	}
	return fin
}

func main() {
	args := ""
	if len(os.Args) > 1 {
		// fmt.Println("")
		args = os.Args[1]
	}
	fmt.Println(alphaMirror(args))

}