package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("XVIIIXLCM"))
}
func romanToInt(s string) int {
    ret := 0
    // is it more performant to use make() syntax and pass capacity to prevent
    // mem reallocations or is it more performant to pass the initial map via
    // literal syntax so map is created at compile time?
    subtractionValuesMap := map[string]bool{
        "IV": true,
        "IX": true,
        "XL": true,
        "XC": true,
        "CD": true,
        "CM": true}
    valuesMap := map[byte]int{
        'I':1,
        'V':5,
        'X':10,
        'L':50,
        'C':100,
        'D':500,
        'M':1000}
    for idx:=0; idx < len(s)-1; idx++ {
        twoChar := s[idx:idx+2]
        if _,ok := subtractionValuesMap[twoChar]; ok {
            ret -= valuesMap[byte(twoChar[0])]
        } else {
            ret += valuesMap[byte(twoChar[0])]
        }
    }
    ret += valuesMap[byte(s[len(s)-1])]
    return ret
}
