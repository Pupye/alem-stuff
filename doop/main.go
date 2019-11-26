package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	length := getLenIsSortedFDoop(arguments)
	if length != 3 {

	} else {
		argType := isValid(arguments[1])
		if argType == -1 {
			fmt.Print(0)
			fmt.Print('\n')
		} else if argType == 1 {
			if arguments[0] == "0" && arguments[2] == "0" {
				fmt.Print(0)
				fmt.Print('\n')
			} else if arguments[0] == "0" {

				fmt.Print(atoiForDoop(arguments[2])) // handle error
				fmt.Print('\n')
			} else if arguments[2] == "0" {
				fmt.Print(atoiForDoop(arguments[2])) //handle error
				fmt.Print('\n')
			} else {

			}

		} else if argType == 2 {

		} else if argType == 3 {

		} else if argType == 4 {
			if arguments[2] == "0" {
				fmt.Print("No division by 0")
				fmt.Print('\n')
			}
		} else if argType == 5 {
			if arguments[2] == "0" {
				fmt.Print("No Modulo by 0")
				fmt.Print('\n')
			}
		}
	}
}

func isValid(a string) int {
	if a == "+" {
		return 1
	} else if a == "-" {
		return 2
	} else if a == "*" {
		return 3
	} else if a == "/" {
		return 4
	} else if a == "%" {
		return 5
	}
	return -1
}

func getLenIsSortedFDoop(tab []string) int {
	isEmpty := true
	length := 0
	for index := range tab {
		length = index
		isEmpty = false
	}
	if !isEmpty {
		length++
	}
	return length
}

func atoiForDoop(s string) (int, int) {
	if s == "" {
		return 0, 0
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
		if mutableOne[index] == 45 && index == 0 {
			return res * (-1), 0
		} else if mutableOne[index] == 43 && index == 0 {
			return res, 0
		}
		if mutableOne[index] < 48 || mutableOne[index] > 57 {
			return 0, 0
		}
		slicedInt := mutableOne[index] - 48
		res += int(slicedInt) * tens
		tens *= 10
	}

	return res, 0
}
