package main

import "fmt"

func main() {

    // variable in for loop is block scoped
    for x := 5; x >= 1; x-- {
        fmt.Printf("%d\n", x)
    }

    // can also use for keyword for while loops
    fX := 0
    for fX < 5 {
        fmt.Println(fX)
        fX++
    }

    arr := []string{"one", "two", "three"}

    for i, str := range arr {
     fmt.Println(i, str)
    }
}
