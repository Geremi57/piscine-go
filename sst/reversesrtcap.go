package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1]
	runes := []rune(args)
	for i := 0; i < len(runes); i++{
		if runes[i] >= 'A' && runes[i] <= 'Z'{
			runes[i] += 32
		}

		if runes[i] >= 'a' && runes[i] <= 'z' && 
		i+1 == len(runes) ||  runes[i +1] == ' ' {
			runes[i] -=32
		} 
	}

	fmt.Println(runes)
	for _,v := range runes{
		z01.PrintRune(v)
		
	}
	z01.PrintRune('\n')
	
	}
	
