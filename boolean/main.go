package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < getLenOfargRune(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	}
	return false

}

func main() {
	arguments := os.Args[1:]

	lengthOfArg := getLenOfarg(arguments)
	if isEven(lengthOfArg) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}

func getLenOfarg(args []string) int {
	isEmpty := true
	length := 0
	for index := range args {
		isEmpty = false
		length = index
	}
	if !isEmpty {
		length++
		return length
	}
	return length
}

func getLenOfargRune(args []rune) int {
	isEmpty := true
	length := 0
	for index := range args {
		isEmpty = false
		length = index
	}
	if isEmpty {
		length++
		return length
	}
	return length
}
