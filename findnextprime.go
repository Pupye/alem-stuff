package piscine

//FindNextPrime is a function that returns next prime
func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	misPrime := false
	for index := nb; index < 2*nb; index++ {
		misPrime = isPrime(index)
		if misPrime {
			return index
		}
	}
	return -1
}

func isPrime(nb int) bool {
	if nb <= 1 {
		return false
	}

	index := 2

	for index*index <= nb {
		if nb%index == 0 {
			return false
		}
		index++
	}

	return true

}
