package piscine

import "github.com/01-edu/z01"

//IsNegative is function that return boolean
func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune(84)
		z01.PrintRune(10)
	} else {
		z01.PrintRune(70)
		z01.PrintRune(10)
	}
}
