package main

import (
	"github.com/01-edu/z01"
)

func main() {
	var alph string = "abcdefghijklmnopqrstuvwxyz\n"
	for i := 0; i < 10; i++ {
		z01.PrintRune(rune(alph[i]))
	}
}