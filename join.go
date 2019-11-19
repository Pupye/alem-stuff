package piscine

//Join is a function
func Join(strs []string, sep string) string {
	build := ""
	length := getLenIsNPQJoin(strs)
	for index := 0; index < length-1; index++ {
		build += strs[index] + sep
	}
	build += strs[length-1]
	return build
}

func getLenIsNPQJoin(s []string) int {

	length := 0
	for x := range s {
		length = x
	}
	length++

	return length
}
