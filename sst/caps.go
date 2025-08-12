package main
import "fmt"

func AlphaCount(s string) string {
	w:=[]rune(s)

	for i:=0; i < len(s); i++{
		if i == 0 || w[i - 1] == ' ' {
			if w[i] >= 'a' && w[i] <= 'z'{
				w[i] = w[i] - 32
			} 
		}else if w[i] >= 'A' && w[i] <= 'Z'{
			w[i] = w[i] + 32
		} 
	}
	final:= ""
	for i:=0; i < len(w); i++{
		final += string(w[i])
	}
	return final
}


func main() {
	fmt.Println(AlphaCount("hello world 72        4455 /"))
}