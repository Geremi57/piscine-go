package main

import "fmt"

func printHex(b byte){
	hex := "0123456789abcdef"
	
	first := b /16
	second := b % 16

	fmt.Print(string(hex[first]))
	fmt.Print(string(hex[second]))

}

func printMemory(arr [10]byte) {
	// runes := []rune{}
	// str := ""
	for i,c := range arr {
		printHex(c)

		if (i+1) % 4 == 0 {
			fmt.Println("yes")
		}else {
			fmt.Print(" ")
		}
	}

	if len(arr)%4 != 0 {
		fmt.Println()
	}

	for _,c := range arr {
		if c < 33 || c > 126 {
			c = '.'
		}
	fmt.Print(string(c))
}
fmt.Println()
		// runes = append(runes, rune(c))
		// if c < 32 || c > 126 {
		// 	fmt.Print(".")
		// }else {
		// 	fmt.Println(string(c))
			
		// }
		// str += string(c)
	}
	// fmt.Println(str)

	// fmt.Println(runes)



func main() {
	printMemory([10]byte{'h','e','l','l','o',16,21,'*'})
	// for _,c := range 
}