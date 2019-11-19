package piscine

//LastRune is a good
func LastRune(s string) rune {
	if s == "" {
		return 0
	}
	runes := []rune(s)
	length := 0
	for x := range runes {
		length = x
	}
	length++
	return runes[length-1]
}
