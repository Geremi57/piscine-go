package format

import (
	"strings"
)

func apostrophes(str string) string {
	// s := strings.Trim(str, " ")
	marks := "'"
	res := []rune{}
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if i > 0 && r == ' ' && strings.ContainsRune(marks, runes[i - 1]){
					continue
					}
		if i > 0 && strings.ContainsRune(marks, r) && res[len(res)-1] == ' '{
				// if str[i - 1] == ' ' {
					res = res[:len(res)-1]
				// }s
			}
			res = append(res, r)
		
		// res = append(res,r)
	}
	// fmt.Println(string(res))

	return string(res)
}