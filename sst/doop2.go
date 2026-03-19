package main

import (
	"fmt"
	"os"
)

func parseInt(s string) (int, bool) {
	if s == "" {
		return 0, false
	}

	sign := 1
	i := 0

	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	if i >= len(s) {
		return 0, false
	}

	result := 0
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}

		digit := int(s[i] - '0')

		// Overflow check before multiplying by 10
		if result > (maxInt-digit)/10 {
			return 0, false
		}

		result = result*10 + digit
	}


	result *= sign

	if result < minInt || result > maxInt {
		return 0, false
	}

	return result, true
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok1 := parseInt(os.Args[1])
	op := os.Args[2]
	b, ok2 := parseInt(os.Args[3])

	if !ok1 || !ok2 {
		return
	}

	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1

	var result int

	switch op {
	case "+":
		if (b > 0 && a > maxInt-b) || (b < 0 && a < minInt-b) {
			return
		}
		result = a + b

	case "-":
		if (b < 0 && a > maxInt+b) || (b > 0 && a < minInt+b) {
			return
		}
		result = a - b

	case "*":
		if a != 0 && (a*b)/a != b {
			return
		}
		result = a * b

	case "/":
		if b == 0 {
			fmt.Println("No division by 0")
			return
		}
		result = a / b

	case "%":
		if b == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		result = a % b

	default:
		return
	}

	fmt.Println(result)
}
