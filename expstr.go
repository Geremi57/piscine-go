package main

import (
	"fmt"
	"os"
)



func main() {
	// fin:=""
	count := 0
	out := ""
	res:=""
	pri:=""

	if len(os.Args) == 2 {
		args := os.Args[1:]

		input := ""

		for i:=0; i < len(args); i++{
			if i > 1 {
				input += ""
			}
			
			input += args[i]
		}
		fmt.Println(input)
		for _,value:= range input {
			if value == ' '{
				count ++
			}else{
				break
			}
			
	}
	for i:=count; i < len(input); i++{
		out += string(input[i])
	}
	fmt.Println(out)

	for i,val := range out {
		if val == ' ' && i > 0 && out[i - 1] != ' '{
			res += "   "
		}
		res += string(val)
	}

	spacecount := 0

	for _,value := range res{
		if value == ' '{
			spacecount ++
			if spacecount <= 3 {
				pri += " "
			}
			}else {
				spacecount = 0
				pri += string(value)
			}	
	}
	fmt.Println(res)
	fmt.Println(pri)
	fmt.Println(count)
	}
}