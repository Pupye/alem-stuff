package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(arrayStr []rune) {

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
		myMessage := []rune{73, 32, 104, 97, 118, 101, 32, 97, 110, 32, 101, 118, 101, 110, 32, 110, 117, 109, 98, 101, 114, 32, 111, 102, 32, 97, 114, 103, 117, 109, 101, 110, 116, 115}

		printStr(myMessage)
	} else {
		mySecond := []rune{73, 32, 104, 97, 118, 101, 32, 97, 110, 32, 111, 100, 100, 32, 110, 117, 109, 98, 101, 114, 32, 111, 102, 32, 97, 114, 103, 117, 109, 101, 110, 116, 115}

		printStr(mySecond)
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
	if !isEmpty {
		length++
		return length
	}
	return length
}
