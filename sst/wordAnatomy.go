package main

import "fmt"

func hasPrefix(str, prefix string) bool{
	if len(prefix) > len(str){
		return false
	}
	for i:=0; i < len(prefix); i++{
		if str[i] != prefix[i]{
			return false
		}
	}
	return true
}

func hasSuffix(str, suffix string) bool{
	if len(suffix) > len(str) {
		return false
	}
	start := len(str) - len(suffix)
	for i:=0; i < len(suffix); i++{
		if str[start+i] != suffix[i]{
			return false
		}
	}
	return true
}

func main() {
	prefixes := []string{"up", "sub", "fr", "vic"}
	suffixes := []string{"er", "scribe", "tory", "gic"}

	word:= "subscribe"
	foundPrefix := ""
	foundSuffix := ""

	for _, p := range prefixes{
		if hasPrefix(word, p){
			foundPrefix += p
			break
		}
	}
	for _,v := range suffixes{
		if hasSuffix(word, v){
			foundSuffix += v
			break
		}
	}
	fmt.Println("prefix:", foundPrefix)
	fmt.Println("suffix:", foundSuffix)

	// wordAnatomy(prefixes, suffixes, word)
}