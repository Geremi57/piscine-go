package main

import "fmt"

func main() {

	//STRNG manipuation
	aString := "Hello"
	aStringChangeable := []byte(aString)
	aStringChangeable[0] = 'A'
	aStringChangeable[1] = 'A'
	aStringChangeable[2] = 'A'
	aStringChangeable[3] = 'A'

	aStringFinal := string(aStringChangeable)
	
	fmt.Println(aString)
	fmt.Println(aStringChangeable)
	fmt.Println(aStringFinal)

//for range loop

slice := []string {
					"word1",
					"word2",
					"word3",
					"word7",
}

for  index,word := range slice {
	fmt.Printf("The index is %v and the value is %v\n", index,word)
	
}
}