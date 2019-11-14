package piscine

//UltimateDivMod is function that return remainder and div of a and b but stores in pointers
func UltimateDivMod(a *int, b *int) {
	resOfDiv := *a / *b
	resOfmod := *a % *b
	*a = resOfDiv
	*b = resOfmod
}
