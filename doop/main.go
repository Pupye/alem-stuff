package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]                 //getting here arguments
	length := getLenIsSortedFDoop(arguments) //getting length of the argument
	if length != 3 {                         // checking whether the argument is only 3
		return
	}

	operatorType := getOperatorType(arguments[1])
	argValid := isValidValues(arguments[0], arguments[2])

	if operatorType == -1 || !argValid {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	} else if operatorType == 1 {
		//here we sure that args are valid
		n1, isOverFlow1 := atoiForDoop(arguments[0])
		n2, isOverFlow2 := atoiForDoop(arguments[2])
		if isOverFlow1 || isOverFlow2 {
		} else {

			if operationOverFlow(n1, n2, 1) {

			} else {

			}
		}
	} else if operatorType == 2 {
		n1, isOverFlow1 := atoiForDoop(arguments[0])
		n2, isOverFlow2 := atoiForDoop(arguments[2])
		if isOverFlow1 || isOverFlow2 {
		} else {

			if operationOverFlow(n1, n2, 2) {

			} else {

			}
		}
	} else if operatorType == 3 {

		n1, isOverFlow1 := atoiForDoop(arguments[0])
		n2, isOverFlow2 := atoiForDoop(arguments[2])
		if isOverFlow1 || isOverFlow2 {

		} else {

			if operationOverFlow(n1, n2, 3) {

			} else {

			}
		}
	} else if operatorType == 4 {
		if arguments[2] == "0" {
			z01.PrintRune('k')
		} else {
			n1, isOverFlow1 := atoiForDoop(arguments[0])
			n2, isOverFlow2 := atoiForDoop(arguments[2])
			if isOverFlow1 || isOverFlow2 {

			} else {

				if operationOverFlow(n1, n2, 4) {

				} else {

				}
			}
		}
	} else if operatorType == 5 {
		if arguments[2] == "0" {
			z01.PrintRune('k')
		}
		n1, isOverFlow1 := atoiForDoop(arguments[0])
		n2, isOverFlow2 := atoiForDoop(arguments[2])
		if isOverFlow1 || isOverFlow2 {

		} else {

			if operationOverFlow(n1, n2, 1) {

			} else {

			}
		}
	}

}

func isValidValues(n1, n2 string) bool {
	if ssNumeric(n1) && ssNumeric(n2) {
		return true
	}
	return false
}

func getOperatorType(a string) int {
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

func atoiForDoop(s string) (int64, bool) {
	if s == "" {
		return 0, false
	}
	mutableOne := []rune(s)
	strLen := 0
	for s := range mutableOne {
		strLen = s
	}
	strLen++ // now we got len
	tens := 1
	var res int64
	res = 0
	for index := strLen - 1; index >= 0; index-- {
		if index == 0 && mutableOne[index] == 45 {
			return res, false
		}
		if mutableOne[index] < 48 || mutableOne[index] > 57 {
			return 0, false
		}

		if mutableOne[0] == 45 {
			slicedInt := mutableOne[index] - 48
			res -= int64(int(slicedInt) * tens)

			if res/int64(tens) != int64(slicedInt)*-1 {
				return 0, true
			}

			tens *= 10
		} else {
			slicedInt := mutableOne[index] - 48
			res += int64(int(slicedInt) * tens)
			if res/int64(tens) != int64(slicedInt) {
				return 0, true
			}
			tens *= 10
		}

	}

	return res, false
}

func ssNumeric(str string) bool {
	if str == "" {
		return false
	}

	mutableOne := []rune(str)
	for index := 0; index < getLenIsIsIs(str); index++ {
		if mutableOne[0] == 45 && index == 0 {
			continue
		}
		isNumeric := mutableOne[index] > 47 && mutableOne[index] < 58
		if !(isNumeric) {
			return false
		}
	}
	return true
}

func getLenIsIsIs(s string) int {
	if s == "" {
		return 0
	}

	length := 0
	runes := []rune(s)
	for x := range runes {
		length = x
	}
	length++

	return length
}

func operationOverFlow(n1 int64, n2 int64, argType int) bool {
	var intMax int64

	intMax = 9223372036854775807

	if argType == 1 {
		if n1 < 0 && n2 < 0 {
			if n1+n2 >= 0 {
				return true
			}
		}
		if n2 >= 0 {
			if n1 > intMax-n2 {
				return true
			}
		} else {
			if n1 < intMax-n2 {

				return true
			}
		}
	} else if argType == 2 {
		if n1 > 0 && n2 > 0 {
			return false
		}

		if n1 < 0 && n2 < 0 {
			return false
		}

		if n1 < 0 {
			if n1-n2 > 0 {
				return true
			}
		} else if n2 < 0 {
			if n1-n2 < 0 {
				return true
			}
		}

	} else if argType == 3 {
		res := n1 * n2
		if res != 0 && res/n1 != n2 {
			return true
		}
	} else if argType == 4 {
		return false
	} else if argType == 5 {
		return false
	}
	return false
}
