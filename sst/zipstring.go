package main

import ("strconv"
"fmt")

func ZipString(s string) string {
	result := ""
	count := 1
	for i:=1; i < len(s); i++ {
		// if s[i] >= 'A' && s[i] <= 'Z' || s[i] >= 'a' && s[i] <= 'z' {
			if s[i] == s[i-1] {
				count++
			}else {
				result += strconv.Itoa(count) + string(s[i-1])
				count = 1
			}
		}

	// }
	result += strconv.Itoa(count) + string(s[len(s)-1])

	return result
}

func main(){
	fv := "YouuungFellllas   weee mobbbbieeeengg"
	fmt.Println(ZipString(fv))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
}