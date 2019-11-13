package main

import "github.com/01-edu/z01"

func main() {
	var alph string = "abcdefghijklmnopqrstuvwxyz\n"
	for i := 0; i < len(alph); i++ {
		z01.PrintRune(rune(alph[i]))
	}
}
