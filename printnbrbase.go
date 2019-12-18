package piscine

import (
	"github.com/01-edu/z01"
)

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
	length := getLenOfStrBase(base)

	if length < 2 {
		return false
	}
	mutableOne := []rune(base)

	for x := range mutableOne {
		if mutableOne[x] == '+' || mutableOne[x] == '-' {
			return false
		}
		for index := x + 1; index < length; index++ {

			if mutableOne[x] == mutableOne[index] {

				return false
			}
		}
	}

	return true
}

func convertNbrIntoBase(n int, base string) { //validated
	if n < 0 {
		z01.PrintRune('-')
		radix := getLenOfStrBase(base)
		requiredLength := requiredLengthToConvert(n, radix)
		count := requiredLength - 1
		baseDigits := []rune(base)
		for count >= 0 {
			getReminderFrom := n / getPower(radix, count)
			digit := (getReminderFrom % radix)
			if digit < 0 {
				digit *= -1
			}
			z01.PrintRune(baseDigits[digit])
			count--
		}
	} else {
		radix := getLenOfStrBase(base)
		requiredLength := requiredLengthToConvert(n, radix)
		count := requiredLength - 1
		baseDigits := []rune(base)
		for count >= 0 {
			getReminderFrom := n / getPower(radix, count)
			digit := getReminderFrom % radix
			z01.PrintRune(baseDigits[digit])
			count--
		}
	}
}

func getPower(nb int, power int) int {
	result := 1
	for index := 0; index < power; index++ {
		result *= nb
	}
	return result
}

func requiredLengthToConvert(n int, b int) int { //validated
	if n == 0 {
		return 1
	}
	length := 0
	if n > 0 {
		for n > 0 {
			length++
			n /= b
		}
	} else {
		for n < 0 {
			length++
			n /= b
		}
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
