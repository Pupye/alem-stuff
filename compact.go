package piscine

//Compact this
func Compact(ptr *[]string) int {
	count := 0
	for index := range *ptr {
		if (*ptr)[index] != "" {
			count++
		}
	}
	compacted := make([]string, count)
	compactIndex := 0
	for index := range *ptr {
		if (*ptr)[index] != "" {
			compacted[compactIndex] = (*ptr)[index]
			compactIndex++
		}
	}
	*ptr = compacted
	return count
}
