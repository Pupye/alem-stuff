package main

import "github.com/01-edu/z01"

func main() {
	var i rune = 122
	for ; i > 96; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune(10)
}