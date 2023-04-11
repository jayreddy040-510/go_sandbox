package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)


func main() {
    rand.Seed(time.Now().Unix())
    randNum := rand.Intn(10) + 1
    fmt.Println(randNum)
    userInputAsInt := takeUserInputReturnInt()
    for userInputAsInt != randNum {
        if userInputAsInt > randNum {
            fmt.Println("Your guess is greater than the number!")
        } else if userInputAsInt < randNum {
            fmt.Println("Your guess is less than the number!")
        }
        userInputAsInt = takeUserInputReturnInt()
    }
    fmt.Println("Your guess is correct!")
    return
}

func takeUserInputReturnInt() int {
    newReader := bufio.NewReader(os.Stdin) // NewReader() def arg is os.Stdin
    userInput, err := newReader.ReadString('\n') // ReadString() def arg is '\n'
    if err == nil {
        userInput = strings.TrimSpace(userInput)
        userInputAsInt, convErr := strconv.Atoi(userInput)
        if convErr != nil {
            return -1
        }
        return userInputAsInt
    } else {
        return -1
    }
}
