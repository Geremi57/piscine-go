package main

import (
	"fmt"
	"os"
)

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	word := os.Args[1]
	firstVowel := -1

	for i, r := range word {
		if isVowel(r) {
			firstVowel = i
			break
		}
	}

	if firstVowel == -1 {
		fmt.Println("No vowels")
		return
	}

	if firstVowel == 0 {
		fmt.Println(word + "ay")
		return
	}

	result := word[firstVowel:] + word[:firstVowel] + "ay"
	fmt.Println(result)
}