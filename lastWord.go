package main

import (
	"fmt"
)

func lastWord(str string) string {
	// fin:= ""
	// result := []string{}
	// for i,v := range str{
	// 	result = append(result, string(str[i]))
	// 	// result = append(result, ",")
	// 	if v == ' '{
	// 		continue
	// 		// fin += string(str[i-1])
	// 	}

	// }
	// for i:=len(result) - 1; i > 0; i-- {
	// 	fin += result[i]
	// 	break
	// }
	// return fin
	end := len(str) - 1

	for end >= 0 && str[end] == ' '{
		end --
	}

if end < 0 {
	return "\n"
}

	start := end

	for start >= 0 && str[start] != ' '{
		start--
	}

	return str[start+1 : end +1] + "\n"
}

func main() {
	fmt.Print(lastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(lastWord("hello data   "))
	// fmt.Print(lastWord("  "))
}
