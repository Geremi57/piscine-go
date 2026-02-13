package main

import "fmt"

func main(){

	fin:=""

	for  i:='9'; i > '0'; i--{
		for j:=i-1; j  > '0'; j--{
			for k:=j-1; k > '0'; k--{
				if i >  j  && j > k{ 
				fin += string(i) + string(j) + string(k) + ", "

}
}
}
}
fmt.Println(fin)
}
