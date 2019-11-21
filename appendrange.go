package piscine

//AppendRange is a good function
func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var s []int
	for index := min; index < max; index++ {
		s = append(s, index)
	}

	return s
}
