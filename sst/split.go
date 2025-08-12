package main
 
import "fmt"

func Split(s, sep string) []string {
	fin:=[]string{}
	current:= ""

	// for i:=0; i < len(s); i++{
		for _,value := range s {
			if string(value) == sep {
				fin = append(fin, current)
				current = ""
				fmt.Println(value)
				// fin[i] = i - 1
			}else {
				current += string(value)
			}
		}
		fmt.Println(current)
	return fin
}

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
// fmt.Println(Split(s, sep string))
}

