package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	for index := range arguments {
		if arguments[index] == "01" || arguments[index] == "galaxy" || arguments[index] == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
