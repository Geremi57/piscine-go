package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arguments,_ := os.ReadFile("standard.txt")
	lines := strings.Split(string(arguments), "\n")


	word := "HELLO"

	for row := 0; row < 8; row++{
		for _,ch := range word {
			start := (int(ch) - 32)*9
			fmt.Print(lines[start+row])
		}
		fmt.Println()
	}

	
	// fmt.Println(lines)
}

