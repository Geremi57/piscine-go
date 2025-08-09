// checkpoint 1 question 2
package main

import (
	"fmt"
)

func retainhalf(str string) {
	var a string = "abcdefghi"
	runes:=[]rune(a)
	// numb:='2'
		fmt.Println(runes)
	l:=len(runes)
	fmt.Print(l)
	f:=[]rune{}
		for _,value:=range runes{
			f=append(f,value)
			fmt.Println(len(f))
		}
		fl:=len(f)
		d:=fl / 2
		for i:=0;i< d;i++{
			fmt.Println(string(f[i]))
			k:=f[i]
			fin:=[]rune{}

				f=append(fin,k)
		}
}

func main() {
		var a string = "abcdefghi"
	retainhalf(a)

}

//iteration -uses a for looop
// gos only loop

func fibonnaci(n int) int{
	if n <= 1{
		return n
	}
	numb:= 0
	onr:=1
	for i:=2; i <= n; i++{
		next := numb + onr
		onr = next
	}
	return numb
}