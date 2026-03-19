package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	arg1 := arguments[0]
	sign:= arguments[1]
	arg2 := arguments[2]
	fmt.Println(arguments)
	n1:=0
	n2:=0
	for _,value := range arg1 {
		n1 = n1 * 10 + int(value - '0')
	}
	for _,value := range arg2 {
		n2 = n2 * 10 + int(value - '0')
	}

	final:=0
	
		// if sign == "+"{
		// 	final = n1 + n2
		// }else if sign == "-" {
		// 	final = n1 - n2
		// }else if sign == "*" {
		// 	final = n1 * n2
		// }else if sign == "/" {
		// 	if n2 == 0{
		// 		fmt.Println("No division by 0")
		// 		return
		// 	}
		// 	final = n1 / n2
		// }else if sign == "%" {
		// 	if n2 == 0{
		// 		fmt.Println("No modulo by 0")
		// 		return
		// 	}
		// 	final = n1 % n2
		// }

		switch sign{
		case "+":
			final = n1 + n2
		case "-":
			final = n1 - n2
		case "*":
			final = n1 * n2
		case "/":
			final = n1 / n2
			if n2 == 0 {
				fmt.Println("no division by 0")
				return
			}
			final = n1 / n2
		case "%":
			if n2 == 0 {
				fmt.Println("no modulo by 0")
				return
			}
			final = n1 % n2
		default:
			return
		
		}

	
	fmt.Println(final)
	// fmt.Print(final)

	// fmt.Print(n1)
	// fmt.Print(n2)


}