package main

import "fmt"

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}

	count := 0
	fin := ""
	skip := false

	for _, v := range str {
		if v == ' ' {
			continue
		}

		if skip {
			skip = false
			continue
		}

		fin += string(v)
		count++

		if count == 5 {
			fin += " "
			count = 0
			skip = true
		}
	}

	if len(fin) < 5 {
		return "Invalid Input\n"
	}

	// if fin[len(fin)-1] == ' ' {
	// 	fin = fin[:len(fin)-1]
	// }

	return fin + "\n"
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}