package main

import "github.com/01-edu/z01"

func main() { 

		var aRune string ="abcdefghijklmnopqrstuvwxyz"

		for i:=0;i<=len(aRune);i++ {
		z01.PrintRune(rune(aRune[i]))
		}
	// fmt.Println("abcdefghijklmnop") 
}