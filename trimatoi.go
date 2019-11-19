package piscine

//TrimAtoi is a function that returns integer from string
func TrimAtoi(s string) int {
	mutableOne := []byte(s)
	for x := range mutableOne {
		if mutableOne[x] < 58 && mutableOne[x] > 47 {
			return getInteger(s) * getSign(s)
		}
	}
	return 0
}

func getInteger(s string) int {
	length := 0
	mutableOne := []byte(s)
	res := 0
	for x := range mutableOne {
		length = x
	}
	length++ // this is the length of a string
	tens := 1
	for index := length - 1; index > -1; index-- {

		if mutableOne[index] < 58 && mutableOne[index] > 47 {
			res = res + tens*int((mutableOne[index]-48))
			tens = tens * 10
		}
	}
	return res
}

func getSign(s string) int {
	mutableOne := []byte(s)
	for x := range mutableOne {
		if mutableOne[x] < 58 && mutableOne[x] > 47 {
			return 1
		}
		if mutableOne[x] == 45 {
			return -1
		}
	}
	return 1
}
