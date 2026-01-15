// package format

// func apostrophes(str string) string {
// 	// s := strings.Trim(str, " ")
// 	// marks := "'"
// 	// res := []rune{}
// 	runes := []rune(str)
// 	inApos := false
// 	end := -1

// 	for i,r := range runes{
// 		// r := runes[i]

// 		//find an apostrophe
// 		if r == '\'' {
// 			if i > 0 && i + 1 < len(runes) && isAlpha(runes[i-1]) && isAlpha(runes[i+1]){
// 				// res = append(res, r)
// 				continue
// 			}
// 			inApos = !inApos
// 		}

// 		if isAlpha(r) {
// 			end = i
// 		}
// 	}
// 			if inApos && end != -1 {
// 				runes = append(runes[:end+1],
// 				append([]rune{'\''}, runes[end + 1:]...)...)
// 				// inApos = true
// 				// start = len(res)
// 				// continue
// 			}

// 			// if inApos && len(res) > 0 && res[len(res)-1] == ' ' {
// 			// 		res = res[:len(res)-1]

// 			// }

// 		// 	inApos = !inApos
// 		// 	// if inApos {
// 		// 		// start = len(res)
// 		// 	// }

// 		// 	res = append(res, r)
// 		// 	continue
// 		// }

// 		// if inApos && r == ' ' && len(res) > 0 && res[len(res)-1] == '\'' {
// 		// 	continue
// 		// }

// 		// if inApos && r == ' ' {
// 		// 	res = append(res, '\'')
// 		// 	inApos = false
// 		// }

// 			// res = append(res, r)

// 		// res = append(res,r)
// 	// }

// 	// if inApos {
// 	// 	res = append(res, '\'')
// 	// }
// 	// fmt.Println(string(res))

// 	return string(runes)
// }

// func isAlpha(r rune) bool {
// 	return (r >= 'a' && r <= 'z' || (r >= 'A' && r <= 'Z'))
// }

package format

func apostrophes(str string) string {
	runes := []rune(str)
	res := []rune{}

	inApos := false

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		// apostrophe found
		if r == '\'' {
			// ignore contractions like don't
			if i > 0 && i+1 < len(runes) &&
				isAlpha(runes[i-1]) && isAlpha(runes[i+1]) {
				res = append(res, r)
				continue
			}

			inApos = !inApos
			res = append(res, r)
			continue
		}

		// word ends → close apostrophe
		if inApos && !isAlpha(r) {
			res = append(res, '\'')
			inApos = false
		}

		res = append(res, r)
	}

	// safety: close if string ended mid-word
	if inApos {
		res = append(res, '\'')
	}

	return string(res)
}

func isAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}
