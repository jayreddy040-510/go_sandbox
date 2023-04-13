package main

import "fmt"

func main () {
    aNums := []int{1, 2, 3}
    for _, val := range aNums {
        fmt.Println(val)
    }

    var aRay [5]int
    fmt.Println(aRay[0], aRay[4])
}
