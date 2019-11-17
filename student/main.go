package student

import (
	"github.com/01-edu/z01"
)

//Raid1a is a function good
func Raid1a(height rune, width rune) {
	if width == 0 || height == 0 {
		//print nothing
	} else if width == 1 && height == 1 {
		z01.PrintRune('o')
		z01.PrintRune('\n')
	} else if width == 1 || height == 1 {
		printHorizont(width, rune('-'), rune('o'), rune('o'))
		printVertical(height, rune('|'), rune('o'), rune('o'))
	} else {
		printHorizont(width, rune('-'), rune('o'), rune('o'))
		var index rune
		for index = 1; index < height-1; index++ {
			printBody(width, rune('|'), rune('|'))
		}
		printHorizont(width, rune('-'), rune('o'), rune('o'))
	}
}

//Raid1b is good function
func Raid1b(height rune, width rune) {
	if width == 0 || height == 0 {
		//print nothing
	} else if width == 1 && height == 1 {
		z01.PrintRune('/')
		z01.PrintRune('\n')
	} else if width == 1 || height == 1 {
		printHorizont(width, rune('*'), rune('/'), rune('\\'))
		printVertical(height, rune('*'), rune('/'), rune('\\'))
	} else {
		printHorizont(width, rune('*'), rune('/'), rune('\\'))
		var index rune
		for index = 1; index < height-1; index++ {
			printBody(width, rune('*'), rune('*'))
		}
		printHorizont(width, rune('*'), rune('\\'), rune('/'))
	}
}

//Raid1c is good function
func Raid1c(height rune, width rune) {
	if width == 0 || height == 0 {
		//print nothing
	} else if width == 1 && height == 1 {
		z01.PrintRune('A')
		z01.PrintRune('\n')
	} else if width == 1 || height == 1 {
		printHorizont(width, rune('B'), rune('A'), rune('A'))
		printVertical(height, rune('B'), rune('A'), rune('C'))
	} else {
		printHorizont(width, rune('B'), rune('A'), rune('A'))
		var index rune
		for index = 1; index < height-1; index++ {
			printBody(width, rune('B'), rune('B'))
		}
		printHorizont(width, rune('B'), rune('C'), rune('C'))
	}
}

//Raid1d is good function
func Raid1d(height rune, width rune) {
	if width == 0 || height == 0 {
		//print nothing
	} else if width == 1 && height == 1 {
		z01.PrintRune('A')
		z01.PrintRune('\n')
	} else if width == 1 || height == 1 {
		printHorizont(width, rune('B'), rune('A'), rune('C'))
		printVertical(height, rune('B'), rune('A'), rune('A'))
	} else {
		printHorizont(width, rune('B'), rune('A'), rune('C'))
		var index rune
		for index = 1; index < height-1; index++ {
			printBody(width, rune('B'), rune('B'))
		}
		printHorizont(width, rune('B'), rune('A'), rune('C'))
	}
}

//Raid1e is good function
func Raid1e(height rune, width rune) {
	if width == 0 || height == 0 {
		//print nothing
	} else if width == 1 && height == 1 {
		z01.PrintRune('A')
		z01.PrintRune('\n')
	} else if width == 1 || height == 1 {
		printHorizont(width, rune('B'), rune('A'), rune('C'))
		printVertical(height, rune('B'), rune('A'), rune('C'))
	} else {
		printHorizont(width, rune('B'), rune('A'), rune('C'))
		var index rune
		for index = 1; index < height-1; index++ {
			printBody(width, rune('B'), rune('B'))
		}
		printHorizont(width, rune('B'), rune('C'), rune('A'))
	}
}

func printHorizont(width rune, signature rune, start rune, end rune) {
	//it should expect integers bigger than 1
	if width > 1 {
		z01.PrintRune(start)
		var index rune = 1
		for ; index < width-1; index++ {
			z01.PrintRune(signature)
		}
		z01.PrintRune(end)
		z01.PrintRune('\n')
	}
}

func printVertical(height rune, signature rune, start rune, end rune) {
	//it should expect integers bigger than 1
	if height > 1 {
		z01.PrintRune(start)
		z01.PrintRune('\n')
		var index rune = 1
		for ; index < height-1; index++ {
			z01.PrintRune(signature)
			z01.PrintRune('\n')
		}
		z01.PrintRune(end)
		z01.PrintRune('\n')
	}
}

func printBody(width rune, start rune, end rune) {
	z01.PrintRune(start)
	var index rune
	for index = 1; index < width-1; index++ {
		z01.PrintRune(' ')
	}
	z01.PrintRune(end)
	z01.PrintRune('\n')
}
