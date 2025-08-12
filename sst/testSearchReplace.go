package main
import ("fmt"
"github.com/01-edu/z01"
"os")

func main() {
	args:= os.Args
	if len(args) < 4 {
		fmt.Println("")
		return
	}
	
	arg1 := args[1]
	arg2 := args[2]
	arg3 := args[3]
	s := []string{}
	for _,value := range arg1 {
		if string(value) == arg2{
			s = append(s, arg3)
		}else {
			s = append(s, string(value))
		}
	}
	final := ""
	for _,value := range s {
		final += string(value)
	}
	fmt.Println(arg1)
	
	fmt.Println(final)

	fmt.Println(s)
	for _,value := range final {
		z01.PrintRune(value)
	}
	// fmt.Println(value)

}