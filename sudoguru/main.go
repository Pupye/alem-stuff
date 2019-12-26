package main

func main() {

}

func transformToInt(s []string, board [][]int) {

	// initialization of board
	for index := 0; index < 9; index++ {
		board[index] = make([]int, 9)
	}

	for i := 1; i < 10; i++ {
		for j, runes := range []rune(s[i]) {
			if runes > '0' && runes <= '9' {
				board[i-1][j] = int(runes - rune(48))
			}
		}
	}
}

func uniqueCol(board [][]int) bool {
	checkingWith := 0

	for indexOfCol := 0; indexOfCol < 9; indexOfCol++ {

		for indexOfRow := 0; indexOfRow < 9; indexOfRow++ {
			if board[indexOfRow][indexOfCol] != 0 {

				checkingWith = board[indexOfRow][indexOfCol]
				for checkIndex := indexOfRow + 1; checkIndex < 9; checkIndex++ {
					if checkingWith == board[checkIndex][indexOfCol] {
						return false
					}
				}
			}
		}
	}
	return true
}

func uniqueRow(board [][]int) bool {
	checkingWith := 0

	for indexOfRow := 0; indexOfRow < 9; indexOfRow++ {

		for indexOfCol := 0; indexOfCol < 9; indexOfCol++ {
			if board[indexOfRow][indexOfCol] != 0 {

				checkingWith = board[indexOfRow][indexOfCol]
				for checkIndex := indexOfCol + 1; checkIndex < 9; checkIndex++ {
					if checkingWith == board[indexOfRow][checkIndex] {
						return false
					}
				}
			}
		}
	}
	return true
}

func uniqueBlocks(board [][]int) {

}
