package piscine

//NRune is a function that returns
func NRune(s string, n int) rune {
	if s == "" {
		return 0
	}
	runes := []rune(s)
	if n < 1 {
		return 0
	}

	length := 0
	for x := range runes {
		length = x
	}
	length++

	if n > length {
		return 0
	}
	return runes[n-1]
}
