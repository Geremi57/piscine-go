package main

import ("os" 
"github.com/01-edu/z01"
)

func main() {
	args:=os.Args

	if len(args) < 3 {
		return 
	}
	
	args1:= args[1]
	args2:= args[2]
	args3:= args[3]
	// fmt.Println(args)
	s:=[]string{}
	// for k:=0; k < len(args); k++{
	for _,value := range args1{

	if string(value) == args2 {
		s = append(s, args3)
		}else{
			s = append(s, string(value))
		}
		
		// s:=make([]string,len(args1))
		// s = append(s[k], s)
	}
	final:=""
	for _,value := range s{
		
		final += value
	}
	// // return s
	// fmt.Println(s)
	
	// fmt.Println(final)
	arr:=[]rune(final)
	for _,value := range arr{
		z01.PrintRune(value)

	}   
	}

	// return final
