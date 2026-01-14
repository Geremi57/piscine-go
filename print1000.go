package main

import (
	"fmt"
)

func main() {

	for a:=0; a <= 7 ;a++{
		for b:=1; b <= 8 ;b++{
			if a >= b {
				continue
			}
			for c:=2; c <= 9 ;c++{
			if b >= c {
				continue
			}
			fmt.Printf("%v%v%v,",a,b,c)
		}
		}	
	}
}
