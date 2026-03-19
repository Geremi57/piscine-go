package main

import "fmt"

func SaveAndMiss(s string, num int) string{
	fin := ""
	arr := []string{}
	for i:=0; i < len(s); i += num {
		end := i + num
		if len(s) < end {
			end = len(s)
		}
		arr = append(arr, s[i:end])

		// if i % 3 == 0 {
		// 	end = i
		// }
		// fin = s[i:end]
	}
	for i,value := range arr{
		if i % 2 == 0{
			fin += value
		} 
	}
	// fmt.Println(arr)
	return fin
}

func main(){
a := "what is your name"
fmt.Println(SaveAndMiss(a, 3))
fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}
