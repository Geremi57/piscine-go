package main

import (
	"fmt"
	"os"
)

func main() {
	// arguments := os.Args[1:]
	
	if len(os.Args) != 2 {
		fmt.Print("")
	}

	s := ""

	start := 0

	for i:=1; i < len(os.Args); i++ {
		if i > 1 {
			s += " "
		}
		s += os.Args[i]
	}
	fmt.Println(os.Args)

	fmt.Println(s)
	for i,value := range s {
		if value != ' ' {
			start = i
			break
		}
	}

	fin := ""
	inspace := false
	for i:= start; i < len(s); i++{
		c := s[i]

		if c != ' '{
			fin += string(c)
			inspace = false
		}else {
			if !inspace {
				fin += "   "
				inspace = true
			}
		}
		
		// fin += s[i]
	}
	fmt.Println(fin)



}