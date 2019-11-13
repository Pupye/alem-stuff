package piscine

import (
	"github.com/01-edu/z01"
)

//PrintComb is function that return boolean
func PrintComb() {
	var i rune = 0
	var j rune = 0
	var k rune = 0
	var count rune = 0
	for i < 10 {
		j = 0
		for j < 10 {
			k = 0
			for k < 10 {
				if i != j && j != k && i != k {
					if i < j && j < k {
						if count == 119 {
							z01.PrintRune(i + 48)
							z01.PrintRune(j + 48)
							z01.PrintRune(k + 48)
							z01.PrintRune(10)
						} else {
							z01.PrintRune(i + 48)
							z01.PrintRune(j + 48)
							z01.PrintRune(k + 48)
							z01.PrintRune(44)
							z01.PrintRune(32)
						}
						count++
					}
				}
				k++
			}
			j++
		}
		i++
	}
}
