package main

import "fmt"

func main() {
	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value
	fmt.Printf("%s %d %c %f %t %o %x\n",
    "string", 3, 'A', 3.14, true, 2, 3)
}
