package piscine

//DivMod is function that return remainder and div result of a and b
func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
}
