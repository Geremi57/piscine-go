package format

import (
	"fmt"
	"strings"
)

func Punctuate(str string) string {
		marks := ".,!?;'"
		runes := []rune(str)
		res := []rune{}
		for i := 0; i < len(runes); i++ {
			r := runes[i]
			// res += string(str[i])
			// if strings.ContainsRune(marks, r) {
				fmt.Println("it contains")
				
				if strings.ContainsRune(marks, r)         {
					if  i > 0 && str[i - 1] == ' '{

						res = res[:len(res)-1]
						// res += string(str[i])
						// if r == ' '{
							fmt.Println(str[i])
							// }
							// fmt.Println("there is space")
						}
					}
				// }
				res = append(res,runes[i])
		}
		fin := ""
		for _,c := range res {
			fin += string(c)
		}
		// fin := strings.Join(string(res), " ")
		fmt.Println(res)
		fmt.Println(fin)
		return fin
}

