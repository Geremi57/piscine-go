package main

import "fmt"

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		fmt.Print("/")
		for i:=0; i < x -2; i++ {
			fmt.Print("*")
		}
		fmt.Println("\\")

		for j:=0;j < y-2; j++{
			fmt.Println("*")
		
		}
		if x > 1 {
		for j:=0;j < x-2; j++{
			fmt.Print(" ")
		}
			fmt.Println("*")
	}else{
		fmt.Println( )
	}
	if  y > 1 {
		fmt.Print("\\")
		if x > 1 {
			for j:=0;j < x-2; j++{
				fmt.Print("*")
			}
			fmt.Print("/")
		}else{
			fmt.Println()
		}
	}

	}
}
func main() {

	QuadB(1,5)
}