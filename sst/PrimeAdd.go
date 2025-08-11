package main

import ("fmt"
"github.com/01-edu/z01"
 "os")

 func Atoi(s string) int{
	n := 0
	for _,v:= range s {
		if v < '0' || v > '9'{
			return 0
		}
		n = n*10 + int(v - '0')
	}
	return n
 }

 func isPrime(n int) bool{
	if n < 2{
		return false
	}

	//if a number n is not a prime it mus have a factor pair where a*b = n 
	for i:=2; i*i <= n ; i++{
		if n%i == 0 {
			return false
		}
	}
	return true

 }

func main(){
	arguments := os.Args
	arg1 := arguments[1]
	if len(arguments) != 2 {
		z01.PrintRune('\n')
	}

	numb := Atoi(arg1)
	prime:=0
	for i := 2; i < numb; i++{
		if isPrime(i){
			prime += i
		}

	}
	fmt.Println(prime)
}