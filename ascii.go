package main
import "fmt"

func Ascii(str string) []byte {
	s := []byte{}
	for _,c := range str {

		s = append(s, byte(c))
		// s += string(c)
	}
	fmt.Println(s)
	return s
}

func main() {
 Ascii("hello")
}