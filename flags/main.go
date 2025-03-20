package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || cont(args, "--help") || cont(args, "-h") {
		printHelp()
		return
	}

	var insert string
	var order bool
	var mainStr string

	for _, arg := range args {
		if startsWith(arg, "--insert=") {
			insert = trimPrefix(arg, "--insert=")
		} else if startsWith(arg, "-i=") {
			insert = trimPrefix(arg, "-i=")
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else {
			mainStr = arg
		}
	}

	mainStr += insert

	if order {
		mainStr = orderString(mainStr)
	}

	fmt.Println(mainStr)
}

func cont(slice []string, v string) bool {
	for _, a := range slice {
		if a == v {
			return true
		}
	}
	return false
}

func startsWith(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	return s[:len(prefix)] == prefix
}

func trimPrefix(s, prefix string) string {
	if startsWith(s, prefix) {
		return s[len(prefix):]
	}
	return s
}

func orderString(s string) string {
	runes := []rune(s)
	bubbleSort(runes)
	return string(runes)
}

func bubbleSort(runes []rune) {
	n := len(runes)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
}

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}
