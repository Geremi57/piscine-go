package main

import "fmt"

func printreverse(s string) string {
	runes := []rune(s)
	l:=len(s)
	// for i, value := range s {
		for i:= 0; i < len(runes) / 2; i++{
			
			runes[i] ,runes[l-1-i] = runes[l-1-i] ,runes[i]
		// }
	}
	fmt.Println( string(runes))
	return string(runes)
}
func main() {
	s := "hello"
	printreverse(s)
}