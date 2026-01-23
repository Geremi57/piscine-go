package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	args1 := arguments[0]
	args2 := arguments[1]
	args3 := arguments[len(arguments) - 1]

	fmt.Println(args1)
	fmt.Println(arguments)

	n1 := 0
	n2 := 0
	for _,c := range args1 {
		n1 = n1 * 10 + int(c - '0')
	}

	for _,c := range args3 {
		n2 = n2 * 10 + int(c - '0')
	}

	res := 0
	if n1 >= 0 && n1 <= 9 && n2 >= 0 && n2 <= 9 {
		if args2 == "+" {
			res = n1 + n2

		}else if args2 == "-" {
			res = n1 - n2
		}else if args2 == "*" {
			res = n1 * n2
		}else if args2 == "/" {
			if n2 == 0 {
				fmt.Println("no division by zero")
				return
			}
			res = n1 / n2
			}else if args2 == "%" {
				if n2 == 0 {
					fmt.Println("no modulo by zero")
					return
			}

			res = n1 % n2
		}
	}


	fmt.Println(res)
	// return n
}
