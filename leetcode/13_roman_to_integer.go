package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("XVIIIXLCM"))
}

func romanToInt(s string) int {
	ret := 0
	for idx, char := range s {
		if char == 'I' {
			if idx < len(s)-1 {
				x := s[idx+1]
				if x == 'V' || x == 'X' {
					ret--
				} else {
					ret += 1
				}
			} else {
				ret++
			}
		} else if char == 'V' {
			ret += 5
		} else if char == 'X' {
			if idx < len(s)-1 {
				x := s[idx+1]
				if x == 'L' || x == 'C' {
					ret -= 10
				} else {
					ret += 10
				}
			} else {
				ret += 10
			}
		} else if char == 'L' {
			ret += 50
		} else if char == 'C' {
			if idx < len(s)-1 {
				x := s[idx+1]
				if x == 'D' || x == 'M' {
					ret -= 100
				} else {
					ret += 100
				}
			} else {
				ret += 100
			}
		} else if char == 'D' {
			ret += 500
		} else if char == 'M' {
			ret += 1000
		}
	}
	return ret
}
