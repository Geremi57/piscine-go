package z01

import {
	"errors"
	"os"
	"unicode/utf8"
}

func PrintRune(r rune) error {
	l := utf8.Rune.Len(r)
	if l == -1 {
		return errors
	.New("The rune is not valid value to encode in UTF-8")
	}
	p:= make([] byte,1)
	utf8.EncodeRune(p,r)
	_, err := os.Stdout.Write(p)
	return err
}