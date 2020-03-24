package main

import (
	"fmt"
	"unicode/utf8"
)

var message string = "This is a recursively called function"

func printStr(str string) {
	printStrRecursive(str, 0)
}

func printStrRecursive(str string, pos int) {
	if pos >= utf8.RuneCountInString(str) {
		fmt.Println()
		return
	} else {
		// fmt.Printf("%c", str[pos])	// Verifies it works for readable characters
		fmt.Printf("%x ", str[pos]) // prints hexadecimals
		printStrRecursive(str, pos+1)
	}
}

func main() {
	printStr(message)
}
