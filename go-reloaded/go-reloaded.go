// package main
// import ("fmt"
// "strconv"
// // "slices"
// "strings"
//  "math")

// func hex(str string) string{
// 	s, err := strconv.ParseInt(str, 16, 64)
// 	// for _,c := range str {
// 	// 	s += string(c)
// 	// }
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(s)

// 	return string(s)
// }

// func bin(str string) string {
// 	s,err := strconv.Atoi(str)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var remain int
// 	decimalNum := 0
// 	index := 0
// 	for s != 0 {
// 		remain = s % 10
// 		s = s / 10
// 		decimalNum = decimalNum + remain * int(math.Pow(2, float64(index)))
// 		index++ 
// 	}
// 	decimal := strconv.Itoa(decimalNum)
// 	fmt.Println(decimalNum)
// 	return decimal
// }

// func up(str string) string{
// 	runes := []rune(str)
// 	// s := ""
// 	for i := 0; i < len(str); i++ {
// 		if runes[i] >= 'a' && runes[i] <= 'z' {
// 			runes[i] -= 32
// 			}
// 	}
// 	// for i,c := range runes {
// 	// 	s[i] += string(c)
// 	// }
// 	fmt.Println(string(runes))
// 	return string(runes)
// }

// func low(str string) string{
// 	 runes := []rune(str)

// 	 for i:= 0; i < len(str); i++ {
// 		if runes[i] >= 'A' && runes[i] <= 'Z' {
// 			runes[i] += 32
// 	 }
// 	 }
// 	 fmt.Println(string(runes))
// 	 return string(runes)
// }

// func cap(str string) string{
// 	 res := strings.Fields(str)
// 	 val := ""

// 	 for i:=0; i < len(res); i++{
// 		val += strings.Title(res[i])
// 		val += " "
// 	 }
// 	//  res = strings.Join(s, " ")

// 	 fmt.Println(val)
// 	 return val
// }

// func lown(str string, n int) string{
// 	// runes := []rune(str)
// 	s:= strings.Fields(str)
// 	val := ""
// 	// index := 0 
// 	// for _,c := range str{
// 	// 	s = append(s, string(c))
// 	// 	if c == ' ' {
// 	// 		// s = append(s, ",")
// 	// 		index += 1
// 	// 	}
// 	// }
// 	// runes := []rune(str)

// 	// lastN:= s[len(s)-n:]
// 	if n > len(s) {
// 		n = len(s)
// 	}

// sp := len(s) - n

// for i := sp; i < len(s); i++ {
// 	s[i] = strings.ToLower(s[i])
// }


// // for i:= 0; i < len(s); i++{
// 	// 		fmt.Println(i)
// 	// 		if runes[i] >= 'a' && runes[i] <= 'z' {
// 		// 			runes[i] -= 32
// 		// 		}
// 		// 		i++
// 		// 	}
// 		// // lastN := /
		
		
// 		// // if err != nil {
// 			// 	// 	panic(err)
// 			// // }
			
// 			// for i,_ := range s {
// 				// 	for i > len(lastN){
// 					// 		fmt.Println(i)
// 					// 		if runes[i] >= 'a' && runes[i] <= 'z' {
// 						// 			runes[i] -= 32
// 						// 		}
// 						// 		i--
// 						// 		}
// 						// }
						
// 						val = strings.Join(s, " ")
						
// 	// fmt.Println(s)
// 	fmt.Println(val)

	

// return val
// }

// func upperN (str string, n int) string{
// 	s := strings.Fields(str)
// 	res := ""

// 	if n > len(s) {
// 		n = len(s)
// 	}

// 	sp := len(s) - n

// 	for i := sp; i < len(s); i++ {
// 		s[i] = strings.ToUpper(s[i])
// 	}

// 	res = strings.Join(s, " ")
// 	fmt.Println(res)

// 	return res

// }

// func capN (str string, n int) string {
// 	s := strings.Fields(str)
// 	res := ""

// 	if n > len(s) {
// 		n = len(s)
// 	}

// 	sp := len(s) - n
// 	for i:= sp; i < len(s); i++{
// 		s[i] = strings.Title(s[i])
// 	}
// 	res = strings.Join(s, " ")

// 	fmt.Println(res)
// 	return res
// }

// func Punctuate (str string) string {
// 		marks := ".,!?;'"
// 		runes := []rune(str)
// 		res := []rune{}
// 		for i := 0; i < len(runes); i++ {
// 			r := runes[i]
// 			// res += string(str[i])
// 			// if strings.ContainsRune(marks, r) {
// 				fmt.Println("it contains")
				
// 				if strings.ContainsRune(marks, r)         {
// 					if  i > 0 && str[i - 1] == ' '{

// 						res = res[:len(res)-1]
// 						// res += string(str[i])
// 						// if r == ' '{
// 							fmt.Println(str[i])
// 							// }
// 							// fmt.Println("there is space")
// 						}
// 					}
// 				// }
// 				res = append(res,runes[i])
// 		}
// 		fin := ""
// 		for _,c := range res {
// 			fin += string(c)
// 		}
// 		// fin := strings.Join(string(res), " ")
// 		fmt.Println(res)
// 		fmt.Println(fin)
// 		return marks
// }


// func apostrophes(str string) string {
// 	// s := strings.Trim(str, " ")
// 	marks := "'"
// 	res := []rune{}
// 	runes := []rune(str)
// 	for i := 0; i < len(runes); i++ {
// 		r := runes[i]
// 		if i > 0 && r == ' ' && strings.ContainsRune(marks, runes[i - 1]){
// 					continue
// 					}
// 		if i > 0 && strings.ContainsRune(marks, r){
// 				if str[i - 1] == ' ' {
// 					res = res[:len(res)-1]
// 				}
// 			}
// 			res = append(res,runes[i])
		
// 		// res = append(res,r)
// 	}
// 	fmt.Println(string(res))
	
// 	return marks
// }

// func AorAn(str string) string {
// 	s := strings.Fields(str)
// 	// fmt.Println(s);
// 	isVowel := false
// 	// res := []string{}
// 	vowels := "AEIOUHaeiouh"

// 	for i:=0; i < len(s); i++ {
// 		if s[i] == "A" || s[i] == "a" {
// 			// fmt.Println("gochaa")
// 			if i + 1 >= len(s) {
// 				continue
// 			}
// 			// if s[i + 1]strings.HasPrefix()
// 			for j,c := range s[i+1] {
// 				// r := s[i]
// 				if j == 0 {
// 					for _,v := range vowels {
// 						if c == v {
// 							isVowel = true
// 							// fmt.Println(string(c))
// 							// break
// 						}
// 					}
					
// 					// fmt.Println(isVowel)
// 					if isVowel {
// 						if s[i] == "A" {
// 							s[i] = "An"
// 						}
// 						if s[i] == "a" {
// 							s[i] = "an"
// 						}
// 						// fmt.Println("its a vowel")
						
// 					}
// 				} 
// 				// return
// 			}
// 			// fmt.Println(s[i + 1])
// 		}
// 	}
// 	fin := strings.Join(s, " ")
// 	fmt.Println(fin)
// 	// fmt.Println(s)
// 	// fmt.Println(isVowel)
// 	// fmt.Println(str)
// 	return fin
// }




// func main() {
// 	hex("1E")
// 	bin("10")
// 	up("hexadecimal")
// 	low("HATE")
// 	cap("this is super amazing")
// 	lown("THIS IS SO EXCITING brothers", 3)
// 	upperN("this is super amazing", 3)
// 	capN("this is super amazing", 3)
// 	Punctuate("I was sitting over there, and then BAMM!!")
// 	apostrophes("I am exactly how they describe me: ' awesome '")
// 	AorAn("There it was. A amazing rock!")
// }

