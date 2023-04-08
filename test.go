package main

import(
	"fmt"
	"bufio"
	"os"
	"reflect"
	"strconv"
	"log"
	"strings"
)

var pl = fmt.Println

func main() {
   reader := bufio.NewReader(os.Stdin)
   iAge, err := reader.ReadString('\n')
   if err != nil {
   	log.Fatal(err)
   }
   iAge = strings.TrimSuffix(iAge, "\n")
   iAgeInt, err2 := strconv.Atoi(iAge)
   if err2 != nil {
   	log.Fatal(err2)
   }
   if (iAgeInt >= 1) && (iAgeInt <= 18){
	pl("Important Birthday")
   } else if (iAgeInt == 21) || (iAgeInt == 50) {
	pl("Important Birthday 21 or 50")
   } else {
	pl("not important bday")
   }
   pl(reflect.TypeOf("\n"), reflect.TypeOf('\n'))
}


