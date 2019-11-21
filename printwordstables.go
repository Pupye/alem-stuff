package piscine

import (
	"github.com/01-edu/z01"
)

//PrintWordsTables is a function
func PrintWordsTables(table []string) {
	forPrint := myConcatParams(table)
	mutableOne := []rune(forPrint)

	for index := range mutableOne {
		z01.PrintRune(mutableOne[index])
	}
	z01.PrintRune('\n')
}

func myConcatParams(args []string) string {
	build := ""

	length := getLenAConCat(args)

	for index := 0; index < length; index++ {
		if index == length-1 {
			build += args[index]
		} else {
			build += args[index] + "\n"
		}
	}

	return build

}

func getLenAConCat(str []string) int {
	length := 0
	isEmpty := true
	for index := range str {
		isEmpty = false
		length = index
	}
	if !isEmpty {
		length++
	}

	return length
}
