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
	validValues := isValidValues(arguments[0], arguments[2])

	if operatorType == -1 || !validValues {
		printStrForDoop("0\n")
		return
	}
	printResultOfOperation(arguments[0], arguments[2], operatorType)

}

func printResultOfOperation(value1, value2 string, operationType int) {
	n1, overFlowCheck1 := atoiForDoop(value1)
	n2, overFlowCheck2 := atoiForDoop(value2)
	if overFlowCheck1 || overFlowCheck2 {
		printStrForDoop("0\n")
		return
	}

	operationOverFlowCheck := operationOverFlow(n1, n2, operationType)
	if operationOverFlowCheck {
		printStrForDoop("0\n")
		return
	}

	if operationType == 1 {

		printAsIntAsRune(n1 + n2)
		z01.PrintRune('\n')
	} else if operationType == 2 {
		printAsIntAsRune(n1 - n2)
		z01.PrintRune('\n')
	} else if operationType == 3 {
		printAsIntAsRune(n1 * n2)
		z01.PrintRune('\n')
	} else if operationType == 4 {
		if n2 == 0 {
			printStrForDoop("No division by 0\n")
			return
		}
		printAsIntAsRune(n1 / n2)
		z01.PrintRune('\n')
	} else if operationType == 5 {
		if n2 == 0 {
			printStrForDoop("No Modulo by 0\n")
			return
		}
		printAsIntAsRune(n1 % n2)
		z01.PrintRune('\n')
	}
}

func printStrForDoop(str string) {
	arrayRune := []rune(str)
	for _, s := range arrayRune {
		z01.PrintRune(s)
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
	}
	return false
}

func printAsIntAsRune(param int64) {
	if param == 0 {
		z01.PrintRune('0')
	} else {

		if param < 0 {
			param = param * -1
			z01.PrintRune('-')
		}

		if param < 10 {
			z01.PrintRune(getRune(param))
		} else {
			var decim int64 = 10
			if param > 1000000000000000000 {
				decim = 1000000000000000000
			} else {
				for param >= decim*10 {
					decim *= 10
				}
			}
			isDone := false
			for param >= 10 && !isDone {
				if param%decim == 0 {
					startDigit := param / decim
					z01.PrintRune(getRune(startDigit))
					for decim > 1 {
						z01.PrintRune('0')
						decim = decim / 10
						isDone = true
					}
				} else {
					startDigit := param / decim
					param = param - startDigit*decim
					decim /= 10
					z01.PrintRune(getRune(startDigit))
				}
			}
			if !isDone {
				if decim > 1 {
					for decim > 1 {
						z01.PrintRune('0')
						decim = decim / 10
						isDone = true
					}
				}
				z01.PrintRune(getRune(param))
			}

		}
	}
}

func getRune(number int64) rune {
	var retRes rune
	if number > 9 {
		z01.PrintRune('-')
	} else {
		var index int64 = 0
		for ; index < number; index++ {
			retRes++
		}

	}
	return retRes + 48
}
