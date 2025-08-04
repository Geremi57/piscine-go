package main

import "github.com/01-edu/z01"

func recMaker(x, y int) {
	if x > 0 && y > 0 {
		z01.PrintRune('1')
		if x > 1 {
			for i:=0;i < x - 2; i++{
				z01.PrintRune('2')
			}
			z01.PrintRune('1')
			}else{
				
				z01.PrintRune('\n')
		}
		for i:=0;i < y-2;i++{
			z01.PrintRune('w')
			z01.PrintRune('\n')
			if x > 1 {

				for i:=0; i<x-2;i++{
					z01.PrintRune(' ')
				}


			}
		}
	}
}

func main() {
	recMaker(5, 4)
}