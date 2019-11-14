package piscine

//StrLen function that prints len of string
func StrLen(str string) int {
	mutableOne := []rune(str)
	i := 0
	for s := range mutableOne {
		i = s
	}
	return i + 1
}
