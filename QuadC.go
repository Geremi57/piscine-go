package main

import "fmt"

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		fmt.Print("/")
		if x > 1 {
		for i:=0; i < x - 2; i++ {
			fmt.Print("*")
		}
		fmt.Println("\\")
	}else {
		fmt.Println()
	}
		for i:=0; i < y-2; i++{
			fmt.Print("*")
			if x > 1 {
				for i:=0;i < x-2;i++{
					fmt.Print(" ")
				}
			fmt.Println("*")
			}else{ fmt.Println()}
		}
	if y > 1 {
		fmt.Print("\\")
		if x > 1 {
		for i := 0; i < x -2; i++ {
			fmt.Print("*")
		}
		fmt.Println("/")
		
	}else {
		fmt.Println()
	}
	}
	}
}

func main() {

	QuadB(5,5)
}