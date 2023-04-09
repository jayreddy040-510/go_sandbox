package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	rStr := "abcdefg"
	runeCount := utf8.RuneCountInString(rStr)
	fmt.Println("rune count:", runeCount)
}
