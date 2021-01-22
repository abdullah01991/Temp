package main

import "fmt"

var y1 = 10
var z1 = "Abdullah Almamun"

/*
Declare the VARIABLE with the IDENTIFIER "z1" is of TYPE String and I am assign the value "Abdullah Almamun"
*/

func main() {
	fmt.Println("The value of y1:", y1)
	fmt.Printf("The type y1: %T\n", y1)

	/*
		This is the STATIC programming language, so VARIABLE is DECLARED
		to hold a VALUE of a certain TYPE, It's now a DYNAMIC programming language
	*/
	fmt.Println("The value of z1:", z1)
	fmt.Printf("The type y1: %T\n", z1)
}
