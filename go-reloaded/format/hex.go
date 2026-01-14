package format

import (
	"fmt"
	"strconv"
)

func hex(str string) string{
	s, err := strconv.ParseInt(str, 16, 64)
	// for _,c := range str {
	// 	s += string(c)
	// }
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	return string(s)
}