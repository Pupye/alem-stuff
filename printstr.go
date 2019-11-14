package piscine

import "github.com/01-edu/z01"

//PrintStr function that prints string chars one by one
func PrintStr(str string) {
	mutableStr := []rune(str)
	for s := range str {
		z01.PrintRune(mutableStr[s])
	}
}
