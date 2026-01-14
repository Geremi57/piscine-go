package main

import (
	"fmt"
)

func CamelToSnakeCase(s string) string{
	final := ""
	resp := true
	for i,v := range s {
		if i == len(s) - 1 && v >= 'A' && v <= 'Z' {
			resp = false
		}else if v >= '0' && v <= '9' {
			resp = false
		}else{
				if i != 0 && v >= 'A' && v <= 'Z' {
							final += "_"
						} 
						final += string(v)
					}
		}
		if !resp{
						final = ""
						for _,v := range s{
							final += string(v)
						}
		}
		return final
}

func main(){
	// arg := os.Args[1:]
	// fin := CamelToSnakeCase("camelToSnakeCase")
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
	// fmt.Println(fin)
}
