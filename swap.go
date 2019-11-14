package piscine

//Swap is function that swaps contents of integers
func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
