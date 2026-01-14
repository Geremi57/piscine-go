package main

import "fmt"

func capitalize(s string) string {
	w := []rune(s)
	fmt.Println(w)

	for i:= 0; i < len(s); i++ {
		if i == 0 || w[i-1] == ' ' {
		if w[i] >= 'a' && w[i] <= 'z' {
			w[i] = w[i] - 32
			// w[i] = w[i] - 32
		}
		} else if w[i] >= 'A' && w[i] <= 'Z' {
			w[i] = w[i] + 32
		}
		 
	}

	fmt.Println(w)
	return string(w)
}
func main() {
	s := "zONE oNE hydra"
	fmt.Println(capitalize(s))

}