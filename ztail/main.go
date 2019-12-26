package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	length := getLen(arguments)
	if length < 2 {
		fmt.Printf("tail: option requires an argument -- 'c'\n")
		fmt.Printf("Try 'tail --help' for more information.\n")
		os.Exit(1)
	}
	numStr := arguments[1]
	plusPresent := false
	switch numStr[0] {
	case '+':
		numStr = numStr[1:]
		plusPresent = true
	case '-':
		numStr = numStr[1:]
	}
	tailLength, isErr := atoi(numStr)
	if isErr {
		if plusPresent {
			fmt.Printf("tail: invalid number of bytes: ‘+%s’\n", numStr)
			os.Exit(1)
		} else {
			fmt.Printf("tail: invalid number of bytes: ‘%s’\n", numStr)
			os.Exit(1)
		}
	}

	if tailLength == 0 && !plusPresent {
		os.Exit(0)
	}
	failed := false
	for index := 2; index < length; index++ {
		f, e := os.Open(arguments[index])

		if e != nil {
			fmt.Printf("tail: cannot open '%s' for reading: No such file or directory\n", arguments[index])
			failed = true
			continue
		}
		if index != 2 && !failed {
			fmt.Printf("\n")
		}

		if length-2 > 1 {
			fmt.Printf("==> %s <==\n", arguments[index])
		}

		size := getFileSize(f)
		if plusPresent && size < tailLength {
			printFormatted("", true)
			continue
		}
		offset, readSize := getOffsetAndReadSize(size, tailLength, plusPresent)
		b := makeByteSlice(readSize)
		f.ReadAt(b, offset)
		printFormatted(string(b), true)
		failed = false
		f.Close()
	}

	os.Exit(0)
}

func printFormatted(s string, isLast bool) {
	if isLast {
		fmt.Printf(s)
	} else {
		fmt.Printf("%s\n", s)
	}
}

func getOffsetAndReadSize(size, tailLength int, plusPresent bool) (int64, int) {
	if plusPresent {
		if tailLength == 0 {
			return 0, size
		}
		return int64(tailLength) - 1, (size - tailLength) + 1
	}

	if tailLength > size {
		return 0, size
	}

	return int64(size - tailLength), tailLength
}

func getFileSize(f *os.File) int {
	i, _ := f.Stat()
	return int(i.Size())
}

func makeByteSlice(size int) []byte {
	res := []byte{0}
	for index := 0; index < size; index++ {
		res = append(res, 0)
	}
	return res
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

func atoi(s string) (int, bool) { //should return error
	if s == "" {
		return 0, true
	}

	mutableOne := []rune(s)
	strLen := getLenStr(s)
	tens := 1
	res := 0
	for index := strLen - 1; index >= 0; index-- {
		if mutableOne[index] < 48 || mutableOne[index] > 57 {
			return 0, true
		}
		slicedInt := mutableOne[index] - 48
		res += int(slicedInt) * tens
		tens *= 10
	}
	return res, false
}

func getLenStr(str string) int {
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
