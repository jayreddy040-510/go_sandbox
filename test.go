package main

import(
	"fmt"
	"bufio"
	"os"
	"log"
	"reflect"
	"strconv"
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
	vv5, err := strconv.Atoi("500000")
	pl(vv5, err, reflect.TypeOf(vv5))
	pl(strconv.Itoa(55))
}

