package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	argumets := os.Args[1:]
	length := getLennDis(argumets)
	if length < 1 {
		fmt.Println("File name missing")
	} else if length > 1 {
		fmt.Println("Too many arguments")
	} else {
		content, err := ioutil.ReadFile(argumets[0])
		if err != nil {
			fmt.Printf(err.Error())
		} else {
			fmt.Printf("%s", content)
		}
	}

}

func getLennDis(args []string) int {
	isEmpty := true
	length := 0
	for index := range args {
		length = index
		isEmpty = false
	}
	if !isEmpty {
		length++
	}
	return length
}
