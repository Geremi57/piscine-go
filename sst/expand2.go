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
		// fmt.Print("")
		arguments := os.Args[1:]

		input := ""
		for i:= 0; i < len(arguments); i++{
			if i > 1 {
				input += ""
			}
			input += arguments[i]
		}
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
		spacecount := 0

		for i, value := range out{
			if value == ' ' && i > 0 && out[i - 1] != ' ' {
				res += "   "
			}
			res += string(value)
		}

		for _,value := range res {
			if value == ' '{
				spacecount ++
				if spacecount <= 3{
					pri += " "
				}
			}else {
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