package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args

	lenght := getLenqqhq(arguments)

	for index := lenght - 1; index > 0; index-- {

		mutableOne := []rune(arguments[index])
		for index2 := range mutableOne {
			z01.PrintRune(mutableOne[index2])
		}
		z01.PrintRune(10)
	}

}

func getLenqqhq(arguments []string) int {

	length := 0
	for x := range arguments {
		length = x
	}
	length++

	return length
}
