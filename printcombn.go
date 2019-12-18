package piscine

import (
	"github.com/01-edu/z01"
)

//PrintCombN is a function
func PrintCombN(n int) {
	count := 0
	for q := 0; q < 10; q++ {
		if n == 1 {
			printPrintCombN(n, q)
			if count != 9 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			count++
			continue
		}
		for w := 0; w < 10; w++ {
			if n == 2 {
				if q < w {
					printPrintCombN(n, q, w)
					if count != 44 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					count++
				}
				continue
			}
			for e := 0; e < 10; e++ {
				if n == 3 {
					if q < w && w < e {
						printPrintCombN(n, q, w, e)
						if count != 119 {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
						count++
					}
					continue
				}
				for r := 0; r < 10; r++ {
					if n == 4 {
						if q < w && w < e && e < r {
							printPrintCombN(n, q, w, e, r)
							if count != 209 {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
							count++
						}
						continue
					}
					for t := 0; t < 10; t++ {
						if n == 5 {
							if q < w && w < e && e < r && r < t {
								printPrintCombN(n, q, w, e, r, t)
								if count != 251 {
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
								count++
							}
							continue
						}
						for y := 0; y < 10; y++ {
							if n == 6 {
								if q < w && w < e && e < r && r < t && t < y {
									printPrintCombN(n, q, w, e, r, t, y)
									if count != 209 {
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
									count++
								}
								continue
							}
							for u := 0; u < 10; u++ {
								if n == 7 {
									if q < w && w < e && e < r && r < t && t < y && y < u {
										printPrintCombN(n, q, w, e, r, t, y, u)
										if count != 119 {
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
										count++
									}
									continue
								}
								for i := 0; i < 10; i++ {
									if n == 8 {
										if q < w && w < e && e < r && r < t && t < y && y < u && u < i {
											printPrintCombN(n, q, w, e, r, t, y, u, i)
											if count != 44 {
												z01.PrintRune(',')
												z01.PrintRune(' ')
											}
											count++
										}
										continue
									}
									for o := 0; o < 10; o++ {
										if n == 9 {
											if q < w && w < e && e < r && r < t && t < y && y < u && u < i && i < o {
												printPrintCombN(n, q, w, e, r, t, y, u, i, o)
												if count != 9 {
													z01.PrintRune(',')
													z01.PrintRune(' ')
												} else {
													z01.PrintRune(10)
													return
												}
												count++
											}
											continue
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune(10)
}

func printPrintCombN(n int, a ...int) {
	switch n {
	case 1:
		z01.PrintRune(rune(a[0]) + 48)

	case 2:
		z01.PrintRune(rune(a[0]) + 48)
		z01.PrintRune(rune(a[1]) + 48)

	case 3:
		z01.PrintRune(rune(a[0]) + 48)
		z01.PrintRune(rune(a[1]) + 48)
		z01.PrintRune(rune(a[2]) + 48)

	case 4:
		z01.PrintRune(rune(a[0]) + 48)
		z01.PrintRune(rune(a[1]) + 48)
		z01.PrintRune(rune(a[2]) + 48)
		z01.PrintRune(rune(a[3]) + 48)
	case 5:
		z01.PrintRune(rune(a[0]) + 48)
		z01.PrintRune(rune(a[1]) + 48)
		z01.PrintRune(rune(a[2]) + 48)
		z01.PrintRune(rune(a[3]) + 48)
		z01.PrintRune(rune(a[4]) + 48)
	case 6:
		z01.PrintRune(rune(a[0]) + 48)
		z01.PrintRune(rune(a[1]) + 48)
		z01.PrintRune(rune(a[2]) + 48)
		z01.PrintRune(rune(a[3]) + 48)
		z01.PrintRune(rune(a[4]) + 48)
		z01.PrintRune(rune(a[5]) + 48)
	case 7:
		z01.PrintRune(rune(a[0]) + 48)
		z01.PrintRune(rune(a[1]) + 48)
		z01.PrintRune(rune(a[2]) + 48)
		z01.PrintRune(rune(a[3]) + 48)
		z01.PrintRune(rune(a[4]) + 48)
		z01.PrintRune(rune(a[5]) + 48)
		z01.PrintRune(rune(a[6]) + 48)
	case 8:
		z01.PrintRune(rune(a[0]) + 48)
		z01.PrintRune(rune(a[1]) + 48)
		z01.PrintRune(rune(a[2]) + 48)
		z01.PrintRune(rune(a[3]) + 48)
		z01.PrintRune(rune(a[4]) + 48)
		z01.PrintRune(rune(a[5]) + 48)
		z01.PrintRune(rune(a[6]) + 48)
		z01.PrintRune(rune(a[7]) + 48)
	case 9:
		z01.PrintRune(rune(a[0]) + 48)
		z01.PrintRune(rune(a[1]) + 48)
		z01.PrintRune(rune(a[2]) + 48)
		z01.PrintRune(rune(a[3]) + 48)
		z01.PrintRune(rune(a[4]) + 48)
		z01.PrintRune(rune(a[5]) + 48)
		z01.PrintRune(rune(a[6]) + 48)
		z01.PrintRune(rune(a[7]) + 48)
		z01.PrintRune(rune(a[8]) + 48)
	}
}
