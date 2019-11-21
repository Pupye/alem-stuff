package piscine

//MakeRange is a function
func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	res := make([]int, max-min)

	for index := min; index < max; index++ {
		res[index-min] = index
	}
	return res
}
