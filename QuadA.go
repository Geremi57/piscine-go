package main

import "fmt"

func QuadA(x, y int){
if x > 0 && y > 0 {
		fmt.Print("o")
		if  x > 1 {
		for i:=0; i < x -2;i++{
			fmt.Printf("-")
		}
		fmt.Println("o")
	}else {
		fmt.Println()
	}
		for i:=0; i < y-2; i++{
			fmt.Print("|")
		if x > 1{
			for k:=0; k < x-2; k++{
			fmt.Print(" ")
			}
			fmt.Println("|")
		}else{ fmt.Println()}
	}
	if y > 1 {
fmt.Print("o")
		if x > 1 {
		for i:=0; i < x -2;i++{
			fmt.Print("-")
		}
		fmt.Println("o")
	}
		}else {
			fmt.Println()
		}
}
}

func main() {
	QuadA(1, 5)
}