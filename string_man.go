package main

import (
	"strings"
	"fmt"
)

var pl = fmt.Println

func main() {
	stringVar1 := "A word's word"
	replacer := strings.NewReplacer("word", "alligator")
	stringVar1Replaced := replacer.Replace(stringVar1)
	pl(stringVar1Replaced)
}
