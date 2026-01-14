package format

import (
	"regexp"
	"strconv"
	"strings"
)

func Formatter(str string) string {
	str = normalizeCommands(str)
	s := strings.Fields(str)
	for i:=0; i < len(s); i++ {
if keyWord(s[i]) {
	operation, n := parse(s[i])
	s = bind(s, i, operation, n)
	i--
}			
		
		
	}
	return strings.Join(s, " ")
	
}


// keyWord detects if the token is a command like (up) or (cap,3)
func keyWord(s string) bool {
		return strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")")
}

func normalizeCommands(s string) string {
	re := regexp.MustCompile(`\(\s*([a-zA-Z]+)\s*,\s*(\d+)\s*\)`)
	return re.ReplaceAllString(s, "($1,$2)")
}


// parse extracts the operation and optional number
func parse(c string) (string, int) {
	c = strings.Trim(c, "()")
	p := strings.Split(c, ",")

	operation := p[0]
	n := 1

	if len(p) == 2 {
		if num,err := strconv.Atoi(strings.TrimSpace(p[1])); err == nil  {
			n = num
		}
	}
	return operation,n
}

// bind applies the operation to the previous n words
func bind(s []string, n int, operate string, j int) []string {
	start :=  n - j
	if start < 0 {
		start = 0
	}

	// for i := start; i < n; i++ {
		switch operate {
			case "up", "low", "cap":
			
			for i := start; i < n; i++ {
			switch operate {
			case "up":
				s[i] = strings.ToUpper(s[i])
			case "low":
				s[i] = strings.ToLower(s[i])
			case "cap":
				s[i] = strings.Title(s[i])
			}
		}
		case "hex":
		s[n - 1] = hex(s[n - 1])
		case "bin":
		s[n - 1] = bin(s[n - 1])
		case "punctuate":
		s[n - 1] = Punctuate(s[n - 1])
		case "aoran":
			s[n - 1 ] = AorAn(s[n - 1])
		case "apostrophes":
		s[n - 1] = apostrophes(s[n - 1])
case "upn":
		text := strings.Join(s[start:n], " ")
		text = upperN(text, j)
		copy(s[start:n], strings.Fields(text))

	case "capn":
		text := strings.Join(s[start:n], " ")
		text = capN(text, j)
		copy(s[start:n], strings.Fields(text))

	case "lown":
		text := strings.Join(s[start:n], " ")
		text = lown(text, j)
		copy(s[start:n], strings.Fields(text))
	
	}
	
	// s = strings.Fields(full) 


	// remove the command token itself
	return append(s[:n], s[n + 1:]...)
}