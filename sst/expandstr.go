package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) != 2 {
		fmt.Print("")
		return
	}
	
	s := ""
	
	start := 0
	
	for i:=1; i < len(os.Args); i++{
		if i > 1 {
			s += " "
		}
		s += os.Args[i]
	}
	fmt.Println(len(os.Args))
	fmt.Println(s)

	for i,value := range s{
		if value != ' '{
			start = i
			break
		}
	}

	str := ""
	inSpace := false

	for i := start; i < len(s); i++{
		c := s[i]
		if c != ' '{
			str += string(c)
			inSpace = false
		}else{
			if !inSpace{
				str += "   "
				inSpace = true
			}
		}
	}
	fmt.Println(str)

}
