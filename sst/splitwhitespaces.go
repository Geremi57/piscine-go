package main

import "fmt"

func Split(s, sep string) string {
 final:=""
	for i:=0; i < len(sep); i++{
	final = sep[:i]
 }
 return final

}

func main(){
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}