package format

import (
	"fmt"
	"strings"
)


func AorAn(str string) string {
	s := strings.Fields(str)
	// fmt.Println(s);
	isVowel := false
	// res := []string{}
	vowels := "AEIOUHaeiouh"

	for i:=0; i < len(s); i++ {
		if s[i] == "A" || s[i] == "a" {
			// fmt.Println("gochaa")
			if i + 1 >= len(s) {
				continue
			}
			// if s[i + 1]strings.HasPrefix()
			for j,c := range s[i+1] {
				// r := s[i]
				if j == 0 {
					for _,v := range vowels {
						if c == v {
							isVowel = true
							// fmt.Println(string(c))
							// break
						}
					}
					
					// fmt.Println(isVowel)
					if isVowel {
						if s[i] == "A" {
							s[i] = "An"
						}
						if s[i] == "a" {
							s[i] = "an"
						}
						// fmt.Println("its a vowel")
						
					}
				} 
				// return
			}
			// fmt.Println(s[i + 1])
		}
	}
	fin := strings.Join(s, " ")
	fmt.Println(fin)
	// fmt.Println(s)
	// fmt.Println(isVowel)
	// fmt.Println(str)
	return fin
}
