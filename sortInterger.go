package main

import "fmt"

func sortInt(arr [8]int) {
	for i := 0; i < len(arr); i++ { 
		for k := 0; k < i; k++ {
    if arr[k] > arr[i] {
			arr[i], arr[k] = arr[k], arr[i]
    }
	}
}
fmt.Println(arr)
}
func main() {
	arr := [8]int{9,3,7,1,6,2,8,0}
	sortInt(arr)

}