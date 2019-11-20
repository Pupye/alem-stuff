package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	mutableOne := []rune(arguments[0])
	for index := range mutableOne {
		z01.PrintRune(mutableOne[index])
	}
	z01.PrintRune(10)
}
