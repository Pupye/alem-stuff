package main

import "github.com/01-edu/z01"

//Door is a struct
type Door struct {
	state bool
}

//OPEN is a var
const OPEN bool = true

//CLOSE is a var
const CLOSE bool = false

//PrintStr is a function
func PrintStr(str string) {
	arrayRune := []rune(str)
	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
}

//CloseDoor is a function
func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...\n")
	ptrDoor.state = CLOSE
}

//OpenDoor is a function
func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...\n")
	ptrDoor.state = OPEN
}

//IsDoorOpen is a function
func IsDoorOpen(Door Door) bool {
	PrintStr("is the Door opened ?\n")
	return Door.state == OPEN
}

//IsDoorClose is a function
func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	return ptrDoor.state == CLOSE
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(*door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
