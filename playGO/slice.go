package main

import "fmt"

func slice(a []string, nbr ...int) []string{
	if len(nbr) == 0 {
		return nil
	}

	start := nbr[0]
	end := len(a)

	if len(nbr) > 1 {
		end = nbr[1]
	}

	if start < 0 {
		start = len(a) + start
	}
	if end < 0 {
		start = len(a) + end
	}

	if start < 0 || end > len(a) || start > end{
		return nil
	}
	return a[start:end]
}

func main() {
	a := []string{"a","b","c","d","e","f"}
	fmt.Println(slice(a, 1))
	fmt.Println(slice(a, 2, 4))
	fmt.Println(slice(a, -3))
	fmt.Println(slice(a, 2, -1))

}