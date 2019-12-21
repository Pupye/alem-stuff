package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	if getLen(arguments) < 1 {
		r := os.Stdin
		defer r.Close()
		for {
			s := io.LimitReader(r, 1)
			q, _ := ioutil.ReadAll(s)
			if len(q) == 0 {
				return
			}
			printStr(string(q))
		}
	} else {
		length := getLen(arguments)
		for index := 0; index < length; index++ {
			bytes, err := ioutil.ReadFile(arguments[index])
			if err != nil {
				printStr(err.Error())
			}
			printStr(string(bytes))
			z01.PrintRune('\n')
		}
	}

}

func printStr(str string) {
	mutableStr := []rune(str)
	for s := range str {
		z01.PrintRune(mutableStr[s])
	}
}

func getLen(str []string) int {
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
