package main

import ("fmt" 
"os")


func isPrime(n int) bool{
	if n < 2 {
		return false
	}	

	for i:=2; i*i <= n ; i++ {
		if n%i == 0 {
			return false
		}

	}
	return true
}

func addPrimeSum(n int) int{
count := 0
	for i:=2; i <= n; i++ {
		if isPrime(i) {
			count+=i
		}
	}
	return count
}

func atoi(s string) int{
	n := 0
	for _,v := range s {
		if v < '0' || v > '9' {
			return 0
		}
		n = n*10 + int(v - '0')
	}
	return n
}

func main() {
	arg := atoi(os.Args[1])


	fmt.Println(addPrimeSum(arg))
}


func atoi(s string) int{
	n := 0
	for _,v := range s{
		if v < '0' || v > '9' {
			return 0
		}
		n = n*10 + int(v-'0')
	}
}


func itoa(n int) string{
	res := ""

	for n > 0 {
		d := n%10

		res += string(d + '0') + res

		n /= 10
	}
}



func atoi(s string) int{
	n := 0

	for _,v := range{
		if v < '0' || v > '9'{
			return 0
		}

		n = int(v - '0') + n*10
	}
}

func itoa(n int) string{
	res := ""

	for  n > 0{
		d := n %10
		res = string(d + '0') + res

		n /= 10
	}
}