package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	str := os.Args[1]

	for _, r := range str {
		count := 1

		if r >= 'a' && r <= 'z' {
			count = int(r-'a') + 1
			fmt.Println(count, r)
		} else if r >= 'A' && r <= 'Z' {
			count = int(r-'A') + 1
		}

		for i := 0; i < count; i++ {
			z01.PrintRune(r)
		}
	}

	z01.PrintRune('\n')
}