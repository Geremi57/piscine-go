package main

import (
	"fmt"
	"go-reloaded/format"
	"os"
)


func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error: not enough commands")
		return
	}

	fileIn := os.Args[1]

	content,err := os.ReadFile(fileIn)

	if err != nil {
		fmt.Println("Error",err)
		// panic(err)
		return
	}

	clean := format.Formatter(string(content))

	err = os.WriteFile("result.txt", []byte(clean),0644)

	if err != nil {
		fmt.Println("error writing this:", err)
		return
	}
	fmt.Println("formmatting done")
}