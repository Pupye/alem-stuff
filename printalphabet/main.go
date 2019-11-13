package main

import "github.com/01-edu/z01"

func main() {
	var i rune = 97
	for ; i < 123; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune(10)
}
