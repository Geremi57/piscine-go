package format

import (
	"fmt"
	"strings"
)

func low(str string) string{
	 runes := []rune(str)

	 for i:= 0; i < len(str); i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 32
	 }
	 }
	 fmt.Println(string(runes))
	 return string(runes)
}

func cap(str string) string{
	 res := strings.Fields(str)
	 val := ""

	 for i:=0; i < len(res); i++{
		val += strings.Title(res[i])
		val += " "
	 }
	//  res = strings.Join(s, " ")

	 fmt.Println(val)
	 return val
}

func lown(str string, n int) string{
	// runes := []rune(str)
	s:= strings.Fields(str)
	val := ""
	// index := 0 
	// for _,c := range str{
	// 	s = append(s, string(c))
	// 	if c == ' ' {
	// 		// s = append(s, ",")
	// 		index += 1
	// 	}
	// }
	// runes := []rune(str)

	// lastN:= s[len(s)-n:]
	if n > len(s) {
		n = len(s)
	}

sp := len(s) - n

for i := sp; i < len(s); i++ {
	s[i] = strings.ToLower(s[i])
}


// for i:= 0; i < len(s); i++{
	// 		fmt.Println(i)
	// 		if runes[i] >= 'a' && runes[i] <= 'z' {
		// 			runes[i] -= 32
		// 		}
		// 		i++
		// 	}
		// // lastN := /
		
		
		// // if err != nil {
			// 	// 	panic(err)
			// // }
			
			// for i,_ := range s {
				// 	for i > len(lastN){
					// 		fmt.Println(i)
					// 		if runes[i] >= 'a' && runes[i] <= 'z' {
						// 			runes[i] -= 32
						// 		}
						// 		i--
						// 		}
						// }
						
						val = strings.Join(s, " ")
						
	// fmt.Println(s)
	fmt.Println(val)

	

return val
}

func upperN(str string, n int) string{
	s := strings.Fields(str)
	res := ""

	if n > len(s) {
		n = len(s)
	}

	sp := len(s) - n

	for i := sp; i < len(s); i++ {
		s[i] = strings.ToUpper(s[i])
	}

	res = strings.Join(s, " ")
	fmt.Println(res)

	return res

}

func capN(str string, n int) string {
	s := strings.Fields(str)
	res := ""

	if n > len(s) {
		n = len(s)
	}

	sp := len(s) - n
	for i:= sp; i < len(s); i++{
		s[i] = strings.Title(s[i])
	}
	res = strings.Join(s, " ")

	fmt.Println(res)
	return res
}
