package main

import (
	"os"

	"github.com/01-edu/z01"
)

// for _, r := range name[0][2:]
func main() {
	name := os.Args[0]

	lastSlashIndex := -1
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '/' {
			lastSlashIndex = i
			break
		}
	}

	if lastSlashIndex != -1 {
		name = name[lastSlashIndex+1:]
	}

	for _, r := range name {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
