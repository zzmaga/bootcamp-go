package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

const (
	CLOSE = false
	OPEN  = true
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...\n")
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...\n")
	ptrDoor.state = CLOSE
	return true
}

func IsDoorOpen(Door *Door) bool {
	PrintStr("is the Door opened ?\n")
	return Door.state == OPEN
}

func IsDoorClose(Door *Door) bool {
	PrintStr("is the Door closed ?\n")
	return Door.state == CLOSE
}

func main() {
	door := &Door{state: OPEN}
	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
