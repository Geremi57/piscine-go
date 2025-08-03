package main

import "fmt"

func QuadE(x, y int) {
	if x > 0 && y > 0 {
		fmt.Print("A")
	
	if x > 1 {
		for i:=0; i<x-2; i++{
			fmt.Print("B")
		}
		fmt.Println("A")
	}else {fmt.Println()}

	for i:=0; i < y -2; i++{
		fmt.Print("B")
		if x > 1{
		for i:=0; i< x -2; i++{
			fmt.Print(" ")
		}
		fmt.Println("B")
		}else{fmt.Println()}
	}
	if y > 1{
		fmt.Print("C")
		if x > 1{
			for i:=0;i < x -2;i++{
				fmt.Print("B")
			}
			fmt.Println("A")
		}else{fmt.Println()}
	}
	
	
	}
}

func main() {
	QuadE(5,3)
	QuadE(5,1)
	QuadE(1,1)
	QuadE(1,5)
}
