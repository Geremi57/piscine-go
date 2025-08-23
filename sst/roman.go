package main

import "fmt"

func Decode(roman string) int {
  prev := 0
  curr := 0
  total := 0
  romans := map[rune] int {
    'I' : 1,
    'V' : 5,
    'X': 10,
    'L' : 50,
    'C': 100,
    'D' : 500,
    'M' : 1000,
  }

  for i:= len(roman) - 1; i >= 0; i-- {
    curr = romans[rune(roman[i])]
    
    if curr < prev {
      total -= curr
    }else {
      total += curr
    }
    prev = curr
  }
  fmt.Println(total)
    return total

}

func main() {
	Decode("IVXLMDMX")
}
