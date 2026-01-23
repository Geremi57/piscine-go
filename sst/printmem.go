package main

import "fmt"

func printHex(b byte) {
	// first := 
	// second := 0
	hex := "0123456789abcdef"


	first := b / 16
	second := b % 16

	fmt.Print(string(hex[first]))
	fmt.Print(string((hex[second])))

}

func printMemory(arr [10]byte) {
	str := ""
	for i,c := range arr {
		printHex(c)

		if (i+1)%4 == 0 {
			fmt.Println()
		}else {
			fmt.Print(" ")
		}
	}

	if len(arr)%4 != 0 {
		fmt.Println()
	}

	for _,j := range arr {
		if j <= 33 {
			j = '.'
		}
		str += string(j)
		// fmt.Println(string(j))
	}
	fmt.Println(str)

	return 
}


func main() {
	printMemory([10]byte{'h','e','l','l','o',16,21,'*'})
}