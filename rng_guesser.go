package main

import (
    "fmt"
    "math/rand"
    "bufio"
    "time"
    "os"
    "strconv"
    "log"
)


func main() {
    rand.Seed(time.Now().Unix())
    randNum := rand.Intn(10) + 1
    newReader := bufio.NewReader(os.Stdin) // NewReader() def arg is os.Stdin
    userInput, err := newReader.ReadString() // ReadString() def arg is '\n'
    if err == nil {
        userInputAsInt, convErr := strconv.Atoi(userInput)
        if convErr != nil {
            log.Fatal(convErr)   
        }
    }
}
