package main
 import "fmt"

func Join(strs []string, sep string) string {
	final:= ""
	for i:=0; i <= len(strs); i++ {
		if i == len(strs) - 1{
			final += strs[i]
			return final
		}else if i != len(strs) - 1{
			final += strs[i] + sep
		}
	}
	return final
}
func main() {
	toConcat := []string{"Hello!", " How", " are", " you"," doing"," today?"}
	fmt.Println(Join(toConcat, ":"))
}