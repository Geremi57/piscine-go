package main

import "fmt"

func atoi(s string) int{
	n := 0
	for  _,v := range s {
		if v < '0' || v > '9' {
			return 0
		}
		n = int(v - '0') + n*10
	}
	return n
}

func itoa(n int) string{
	res := ""
	for n > 0 {
		d := n%10
		res = string(d + '0') + res
		n/=10
	}
	return res
}

func NotDecimal(dec string) string {
	fin := ""
	if dec == "" {
		return  "\n"
	}
	dot := -1

	for i := 0; i < len(dec); i++ {
		if dec[i] == '.' {
			dot = i
			break
		}
	}

	if dot == -1 {
		return dec + "\n"
	}

	intP := dec[:dot]
	frac := dec[dot+1:]

	for _,c := range dec{
		if (c < '0' || c > '9') && c != '.' && c != '-' {
			return dec + "\n"
		}
	}
	fin = intP + frac
	// var res int
	// var result string

	// for _,v := range fin {
	// 	if v == '-' {
	// 		return fin + "\n"
	// 	}else{
	// 		res = atoi(fin)
	// 		result = itoa(res)

	// 	}
	// }
	res := ""
	for i,v := range fin{
		if v == '0' {
			if i == 1 {
				res += string(v)
			}
				// res += string(v)

		}else {
			res += string(v)
			// break
		}
				// res += string(v)

	}


	return res+ "\n"
}


func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}