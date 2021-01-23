package main

import "fmt"

var x2 int = 10
var y2 string = "Abdullah"
var z2 bool = true

func main() {
	//Print all variable one by one	 and call global variable
	fmt.Println(x2)
	fmt.Println(y2)
	fmt.Println(z2)
	//Print all variable together and call global variable
	fmt.Printf("================================================\n")
	fmt.Println(x2, y2, z2)
	s := fmt.Sprintf("x = %v & Type:%T \ny = %v & Type:%T \nz = %v & Type:%T \n", x2, x2, y2, y2, z2, z2)
	fmt.Println(s)
	fmt.Printf("================================================\n")
	//Print all variable one by one	 and call local variable
	x2 = 10 + 10
	y2 = "James Bond"
	z2 = false

	fmt.Println(x2)
	fmt.Printf("%T\n", x2)
	fmt.Println(y2)
	fmt.Printf("%T\n", y2)
	fmt.Println(z2)
	fmt.Printf("%T\n", z2)
	fmt.Printf("================================================\n")
	//Print all variable together and call global variable

	x2, y2, z2 = 30, "Abdullah", false

	fmt.Println(x2, y2, z2)
	s = fmt.Sprintf("x = %v & Type:%T \ny = %v & Type:%T \nz = %v & Type:%T \n", x2, x2, y2, y2, z2, z2)
	fmt.Println(s)
}
