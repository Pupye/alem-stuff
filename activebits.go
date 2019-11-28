package piscine

//ActiveBits is a functions
func ActiveBits(n int) uint {
	count := 0
	if n < 0 {
		n = n * -1
	}
	for n > 0 {
		bit := n % 2
		if bit == 1 {
			count++

		}
		n = n / 2
	}
	return uint(count)
}
