package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args

	for index := range arguments {
		if index == 0 {
			continue
		}
		mutableOne := []rune(arguments[index])
		for index2 := range mutableOne {
			z01.PrintRune(mutableOne[index2])
		}
		z01.PrintRune(10)
	}

}
