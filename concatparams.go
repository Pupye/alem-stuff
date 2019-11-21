package piscine

//ConcatParams is a function
func ConcatParams(args []string) string {
	build := ""

	length := getLenConCat(args)

	for index := 0; index < length; index++ {
		if index == length-1 {
			build += args[index]
		} else {
			build += args[index] + "\n"
		}
	}

	return build

}

func getLenConCat(str []string) int {
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
