package main

import (
	"fmt"
	"reflect"
	"strings"
)

var pl = fmt.Println

func main() {
	stringVar1 := "A word's word"
	untrimStr := "\nIAmNotTrimmed\n"
	
	// replacer pattern better for multiple replacements, call strings.Replace() for few replacements
	replacer := strings.NewReplacer("word", "alligator")
	stringVar1Replaced := replacer.Replace(stringVar1)
	pl(stringVar1Replaced)
	
	// string interpolation in golang doesn't print said str, forms the string, call Println() to print
	msg := fmt.Sprintf("length of stringVar1Replaced: %d", len(stringVar1Replaced))
	pl(msg)

	fmt.Println("stringVar1 contains A:", strings.Contains(stringVar1, "A"))
	fmt.Println("index of \"word\" in stringVar1:", strings.Index(stringVar1, "word"))

	// last arg of strings.Replace() indicates how many instances you want to replace, -1 = all
	fmt.Println("replacing stringVar1 \"o\" with \"0\":", strings.Replace(stringVar1, "o", "0", -1))

	fmt.Println("Trim this str:", untrimStr, strings.TrimSpace(untrimStr))

	fmt.Println("stringVar1 lowered:", strings.ToLower(stringVar1))
	fmt.Println("stringVar1 uppered:", strings.ToUpper(stringVar1))

	// prefix/suffix bools
	fmt.Println("Prefix for stringVar1 == A?:", strings.HasPrefix(stringVar1, "A"))
	fmt.Println("Prefix for stringVar1 == B?:", strings.HasPrefix(stringVar1, "B"))
	fmt.Println("Suffix for stringVar1 == word?:", strings.HasSuffix(stringVar1, "word"))
	fmt.Println("Split a-b-c-d  by - delimiter:", reflect.TypeOf(strings.Split("a-b-c-d", "-")))
}
