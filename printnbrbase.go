package piscine

import "github.com/01-edu/z01"

//PrintNbrBase is a function
func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		printStrForBase("NV")
		return
	}
	convertNbrIntoBase(nbr, base)
}

func printStrForBase(str string) { //validated
	mutableStr := []rune(str)
	for s := range str {
		z01.PrintRune(mutableStr[s])
	}
}

func isValidBase(base string) bool {

	return false
}

func convertNbrIntoBase(n int, base string) []rune {
	// radix := getLenOfStrBase(base)
	// requiredLength := requiredLengthToConvert(n, radix)

	return nil
}

func requiredLengthToConvert(n int, b int) int { //validated
	if n < b {
		return 1
	}
	length := 0
	for n > 0 {
		length++
		n /= b
	}
	return length
}

func getLenOfStrBase(str string) int { //validated
	isEmpty := true
	length := 0
	for index := range str {
		isEmpty = false
		length = index
	}
	if !isEmpty {
		length++
	}
	return length
}
