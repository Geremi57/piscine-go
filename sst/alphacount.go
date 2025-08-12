package main

import "fmt"

func AlphaCount(s string) int {
count:=0

for _,value := range s{
	if value >= 'a' && value <= 'z' || value >= 'A' && value <= 'Z' {
		count += 1
	}
}
return count

}

func main() {
	fmt.Println(AlphaCount("Hello 78lplp 90World!    4455 /"))

}