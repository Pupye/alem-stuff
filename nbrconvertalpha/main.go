package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	upperFlag := false
	for index := range arguments {
		if index == 0 {
			continue
		}

		if index == 1 && arguments[index] == "--upper" {
			upperFlag = true
			continue
		}

		if upperFlag {
			z01.PrintRune(getLatingivenString(arguments[index]) - 32)
		} else {
			z01.PrintRune(getLatingivenString(arguments[index]))
		}
	}
	z01.PrintRune(10)
}

func getLatingivenString(s string) rune {
	mutableOne := []rune("abcdefghijklmnopqrstuvwxyz")
	lengthOfAlphs := 26
	index := myAtoi(s)

	if lengthOfAlphs < index {
		return ' '
	}
	return mutableOne[index-1]
}

func myAtoi(s string) int {
	if s == "" {
		return 0
	}
	mutableOne := []rune(s)
	strLen := 0
	for s := range mutableOne {
		strLen = s
	}
	strLen++ // now we got len
	tens := 1
	res := 0
	for index := strLen - 1; index >= 0; index-- {
		notInNumbericRange := mutableOne[index] < 48 || mutableOne[index] > 57
		if notInNumbericRange {
			return 27
		}
		slicedInt := mutableOne[index] - 48
		res += int(slicedInt) * tens
		tens *= 10
	}

	return res
}
