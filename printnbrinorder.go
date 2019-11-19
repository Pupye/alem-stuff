package piscine

import (
	"github.com/01-edu/z01"
)

//PrintNbrInOrder is a function that prints digits in order
func PrintNbrInOrder(n int) {
	if n > -1 {
		var digits []int
		for n >= 10 {
			digit := n % 10
			n = n / 10
			digits = append(digits, digit)
		}
		digits = append(digits, n)
		sortIntegerTable(digits)
		for x := range digits {
			z01.PrintRune(rune(digits[x]) + 48)
		}
	}
}

func sortIntegerTable(table []int) {
	key := 0
	j := 0

	length := 0
	for s := range table {
		length = s
	}
	length++

	for index := 1; index < length; index++ {
		key = table[index]
		j = index - 1

		for j >= 0 && table[j] > key {
			table[j+1] = table[j]
			j = j - 1
		}
		table[j+1] = key
	}
}
