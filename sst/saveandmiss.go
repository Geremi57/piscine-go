package main

import "fmt"

func saveandmiss(s string, num int) string{
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
fmt.Println(saveandmiss(a, 3))
}
