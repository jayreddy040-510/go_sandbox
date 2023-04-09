package main

import (
	"fmt"
	"time"
)

// behaves similarly to datetime in python
func main() {
	now := time.Now()
	fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

}
