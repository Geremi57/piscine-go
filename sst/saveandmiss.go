package main

import "fmt"

func saveandmiss(s string, num int) string{
	arr := []string{}
	
	for i:=0; i < len(s); i += num{
		end := i + num
		if end > len(s) {
			end = len(s)
		}
		arr = append(arr, s[i : end])
	}

	final := ""
	for i,value := range arr{
		if i % 2 == 0 {

			final += value
		}
	}
	return final

}

func main(){
	arg1 := "123456789"
	arg2 := "abcdefghijklmnopqrstuvwxyz"

	fmt.Println(saveandmiss(arg1, 3))
	fmt.Println(saveandmiss(arg2, 3))

}