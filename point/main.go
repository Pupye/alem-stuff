package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printAsRuneInt(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printAsRuneInt(points.y)
	z01.PrintRune(10)
}

func printAsRuneInt(param int) {
	if param == 0 {
		z01.PrintRune('0')
	} else {

		if param < 0 {
			param = param * -1
			z01.PrintRune('-')
		}

		if param < 10 {
			z01.PrintRune(getRune(param))
		} else {
			decim := 10
			for param >= decim*10 {
				decim *= 10
			}
			isDone := false
			for param >= 10 && !isDone {
				if param%decim == 0 {
					startDigit := param / decim
					z01.PrintRune(getRune(startDigit))
					for decim > 1 {
						z01.PrintRune('0')
						decim = decim / 10
						isDone = true
					}
				} else {
					startDigit := param / decim
					param = param - startDigit*decim
					decim /= 10
					z01.PrintRune(getRune(startDigit))
				}
			}
			if !isDone {
				z01.PrintRune(getRune(param))
			}

		}
	}
}

func getRune(number int) rune {
	var retRes rune
	if number > 9 {
		z01.PrintRune('-')
	} else {

		for index := 0; index < number; index++ {
			retRes++
		}

	}
	return retRes + 48
}
