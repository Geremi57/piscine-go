package main

import ("os" 
"fmt")

func main() {
	file,err := os.Create("data.txt")

	if err != nil {
		fmt.Print("Error:",err)
	}
	file.WriteString("zipstring\ncountch\nitoa\naddprim\nnotdec\nfindpa\n")

	defer file.Close()
}