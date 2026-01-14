package format

import (
	"fmt"
	"math"
	"strconv"
)

func bin(str string) string {
	s,err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	var remain int
	decimalNum := 0
	index := 0
	for s != 0 {
		remain = s % 10
		s = s / 10
		decimalNum = decimalNum + remain * int(math.Pow(2, float64(index)))
		index++ 
	}
	decimal := strconv.Itoa(decimalNum)
	fmt.Println(decimalNum)
	return decimal
}

