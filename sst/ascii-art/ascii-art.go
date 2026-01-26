package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arguments,err := os.ReadFile("standard.txt")

	if err != nil {
		return
	}
	lines := strings.Split(string(arguments), "\n")

	word := ""
	
	for i:=1; i < len(os.Args); i++{
		word += os.Args[i]
	}
	word = strings.ReplaceAll(word, `\n`, "\n")
	parts := strings.Split(word, "\n")
	// fmt.Println(word)

	for _,part := range parts {
		for row := 0; row <= 8; row++{
			for _,ch := range part {
				if ch < 32 || ch > 126 {
					continue
				}
		start := (int(ch) - 32)*9
		fmt.Print(lines[start+row])
	}
	fmt.Println()
}
}

	
	// fmt.Println(lines)
}

