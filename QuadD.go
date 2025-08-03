package main

import "fmt"

func QuadD(x, y int) {
	if x > 0 && y > 0 {
		fmt.Print("A")
		if x > 1 {
			for i:=0; i < x - 2; i++{
				fmt.Print("B")
			}
			fmt.Println("C")
		}else {
			fmt.Println()
		}
			for i:=0; i < y-2;i++{
				fmt.Print("B")
				if x > 1 {
					for i:=0; i < x-2; i++{
						fmt.Print(" ")
					}
			fmt.Println("B")
			}else {
				fmt.Println()
			}
		}
		if y > 1 {
			fmt.Print("A")
			if x > 1 {
		for i:=0; i < x-2;i++{
			fmt.Print("B")
		}
		fmt.Println("C")
	}else{fmt.Println()}
		}

	}
}

func main() {
	QuadD(5,3)
	QuadD(5, 1)
	QuadD(1,1)
	QuadD(1,5)
}