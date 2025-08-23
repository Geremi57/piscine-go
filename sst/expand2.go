package main

import ("os"
		"fmt")

func main() {
	fin := ""
	count := 0
	out := ""
	res := ""
	pri:= ""
	if len(os.Args) == 2 {
		arguments := os.Args[1:]


		//join all arguments
		input := ""
		for i:= 0; i < len(arguments); i++{
			if i > 1 {
				input += ""
			}
			input += arguments[i]
		}
		fmt.Println(input)
		//checks how many " " 
		// characters there are before the next 
		// letter or actual character like a number
		for _,value := range input {
			if value == ' '{
				count ++
				}else{
					break
				}
			}

			for i:= count; i < len(input); i++{
				out += string(input[i])
			}
		fmt.Println(out)
		//this will add three spaces between words
			for i, value := range out{
				if value == ' ' && i > 0 && out[i - 1] != ' ' {
					res += "   "
				}
				res += string(value)
			}
		fmt.Println(res)
			
			//checks how many space characters there are if there are any
			spacecount := 0
		for _,value := range res {
			if value == ' '{
				spacecount ++
				if spacecount <= 3{
					pri += " "
				}
			}else {
				// if the value we are in is not ' ' reset the space count to 0
				spacecount = 0
				pri += string(value)
			}
		}
		}
	fmt.Println(res)
	fmt.Println(pri)

	fmt.Println(out)
	fmt.Println(count)
	fmt.Println(fin)
}