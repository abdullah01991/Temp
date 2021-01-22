package main

import "fmt"

var a1 int

type abdullah int

var b1 abdullah

/*
This is the way to careate our own Data Type
Here abdullah is the data type from main package
And b1 is variable for the type abdullah
*/

func main() {

	a1 = 20
	b1 = 30
	fmt.Println(a1)
	fmt.Printf("Type of a1 value: %T\n", a1)
	// Data Print from type Abdullah
	fmt.Println(b1)
	fmt.Printf("Type of b1 value: %T\n", b1)

	/*
	   Conversion of type means we take certain Data Type and Converted is to different
	   Data Type.
	*/

	a1 = int(b1)
	fmt.Println(a1)
	fmt.Printf("Type of a1 value after Conversion: %T\n", a1)

}
