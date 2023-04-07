package main

import(
	"fmt"
	"bufio"
	"os"
	"log"
	"reflect"
)

var pl = fmt.Println

func main() {
	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	};
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
}

