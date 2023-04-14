package main

import(
    "strconv"
    "fmt"
      )

func main() {
    fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    stringOfInt := strconv.Itoa(x)
    i := 0
    j := len(stringOfInt) - 1

    for i < j {
        if stringOfInt[i] != stringOfInt[j] {
            return false
        }
        j--
        i++
    }
    return true
}


