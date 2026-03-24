package main

import (
	"fmt"
)

func CamelToSnakeCase(s string) string{
	fin := ""
	for i,v := range s{
		if v >= 'A' && v <= 'Z'{
			// fin += string(v)
			if i != 0 {
				fin += "_"
			}
			} 
			fin += string(v)
	}
	return fin
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}