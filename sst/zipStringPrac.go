package main

import "fmt"
		
func itoa(n int) string{
	// n := 0

	res := ""
	for n > 0 {
		d := n % 10
		res = string(rune(d) + '0') + res

		n /= 10
	}
	return res

}

func ZipString(s string) string {
	result := ""
	count := 1
	for i:=1; i < len(s); i++{
		if s[i] == s[i-1] {
			count++
		}else{
			result += itoa(count) + string(s[i-1])
			count = 1
		}
	}
			result += itoa(count) + string(s[len(s)-1])

	return  result
}

func main(){
	fv := "YouuungFellllas   weee mobbbbieeeengg"
	fmt.Println(ZipString(fv))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
}
