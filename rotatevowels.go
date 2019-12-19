package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[:1]
	argLen := getLenRotate(arguments)
	if argLen != 1 {
		z01.PrintRune('\n')
		return
	}
	mutableStr := []rune(arguments[0])                // need to convert this string
	positionVowel := makeMapPositionVowel(mutableStr) //map position into vowel
	for position, value := range mutableStr {

	}
}

func makeMapPositionVowel(mutableStr []rune) map[int]rune {
	return nil
}

func getLenRotate(s []string) int {

	length := 0
	for x := range s {
		length = x
	}
	length++

	return length
}
