package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	bubbleSortTable(arguments)

	for x := range arguments {
		if x == 0 {
			continue
		}
		mutableOne := []rune(arguments[x])
		z01.PrintRune(mutableOne[0])
		z01.PrintRune(10)
	}

}

func bubbleSortTable(table []string) {

	for i := 1; i < getLenqqhqSort(table)-1; i++ {

		for j := 1 + i; j < getLenqqhqSort(table); j++ {

			if table[i][0] > table[j][0] {

				table[i], table[j] = table[j], table[i]

			}

		}

	}

}

func getLenqqhqSort(table []string) int {
	length := 0

	for x := range table {
		length = x
	}
	length++

	return length
}
