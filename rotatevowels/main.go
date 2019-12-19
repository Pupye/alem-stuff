package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	argLen := getLenRotate(arguments)
	if argLen < 1 {
		z01.PrintRune('\n')
		return
	}
	mutableStr := []rune(joinRotate(arguments, " ")) // need to convert this string
	start := 0
	end := len(mutableStr)
	for start < end {
		if isVowel(mutableStr[start]) {
			temp := mutableStr[start]
			isVowelAtEnd := false
			for !isVowelAtEnd {
				end--
				isVowelAtEnd = isVowel(mutableStr[end])
				if isVowelAtEnd {
					break
				}
			}
			mutableStr[start] = mutableStr[end]
			mutableStr[end] = temp
		}
		start++
	}
	printStrRotate(mutableStr)
	z01.PrintRune('\n')
}

func getLenRotate(s []string) int {
	length := 0
	isEmpty := true

	for x := range s {
		isEmpty = false
		length = x
	}

	if !isEmpty {
		length++
	}

	return length
}

func isVowel(char rune) bool {
	vowels := []rune("aeiouAEIOU")
	for _, v := range vowels {
		if char == v {
			return true
		}
	}
	return false
}

func joinRotate(strs []string, sep string) string {
	build := ""
	length := getLenRotate(strs)
	for index := 0; index < length-1; index++ {
		build += strs[index] + sep
	}
	build += strs[length-1]
	return build
}

func printStrRotate(mutableStr []rune) { //validated
	for s := range mutableStr {
		z01.PrintRune(mutableStr[s])
	}
}
