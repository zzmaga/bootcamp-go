package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printError(err error) {
	for _, r := range "ERROR: " {
		z01.PrintRune(r)
	}
	for _, r := range err.Error() {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) <= 1 {
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil && err != io.EOF {
				printError(err)
				os.Exit(1)
			}
			if n == 0 {
				break
			}

			for i := 0; i < n; i++ {
				z01.PrintRune(rune(buf[i]))
			}
			if err == io.EOF {
				break
			}
		}
		return
	}

	for _, fileName := range os.Args[1:] {
		file, err := os.Open(fileName)
		if err != nil {
			printError(err)
			os.Exit(1)
		}
		defer file.Close()

		buf := make([]byte, 1024)
		for {
			n, err := file.Read(buf)
			if err != nil && err != io.EOF {
				printError(err)
				os.Exit(1)
			}
			if n == 0 {
				break
			}

			for i := 0; i < n; i++ {
				z01.PrintRune(rune(buf[i]))
			}
			if err == io.EOF {
				break
			}
		}
	}
}
