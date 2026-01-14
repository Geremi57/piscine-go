<<<<<<< HEAD
package main

import "fmt"

func RevConcatAlternate(slice1, slice2 []int) []int {
	fin := []int{}
	res := []int{}
	minlength:= len(slice1)

	if minlength > len(slice2) {
		minlength = len(slice2)
	}

	for i := 0; i < minlength; i++ {
		fin = append(fin, slice1[i], slice2[i])
	}
	// return fin
	if minlength < len(slice1) {
		fin = append(fin, slice1[minlength:]...)
	}
	if minlength < len(slice2) {
		fin = append(fin, slice2[minlength:]...)
	}
	
	for i:=len(fin) - 1; i >= 0; i--{
		res = append(res, fin[i])
	}
	return res

}

func main() {
	fmt.Println(RevConcatAlternate([]int{1,2,3}, []int{4,5,6,7,8}))
}
=======
package main

import "fmt"

func RevConcatAlternate(slice1,slice2 []int) []int{
slice1b := slice1
slice2b := slice2
slice3 := []int{}
l:= len(slice1b)
l2:= len(slice1b)


for i:=0; i < len(slice1b) / 2; i++{
	slice1b[l - 1 -i], slice1b[i] = slice1b[i], slice1b[l - 1 - i]
}
for i:=0; i < len(slice2b) / 2; i++{
	slice2b[l2 - 1 -i], slice2b[i] = slice2b[i], slice2b[l2 - 1 - i]
}
fmt.Println(slice2b)
fmt.Println(slice1b)

la := l + l2
i ,k := 0, 0

if l >= l2{
	slice3 = append(slice3, slice1b[0])
	i = 1
	k = 0
}else {
	slice3 = append(slice3, slice2b[0])
	k = 1
	i = 0
}
for  len(slice3) < la {
	if i < len(slice1b){
		slice3 = append(slice3, slice1b[i])
		i++
	}
	 if k < len(slice2b){
		slice3 = append(slice3, slice2b[k])
		k++
	}
}

return slice3

}

func main(){
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
}
>>>>>>> 5541715 (write md for first question)
