package main

import "fmt"

var x1 int
var y1 string
var z1 bool

func main() {
	//Print all variable one by one	 and call global variable
	fmt.Println(x1)
	fmt.Println(y1)
	fmt.Println(z1)
	//Print all variable together and call global variable
	fmt.Printf("================================================\n")
	fmt.Println(x1, y1, z1)
	fmt.Printf("x = %v & Type:%T \ny = %v & Type:%T \nz = %v & Type:%T \n", x1, x1, y1, y1, z1, z1)
	fmt.Printf("================================================\n")
	//Print all variable one by one	 and call local variable
	x1 = 10
	y1 = "James Bond"
	z1 = true

	fmt.Println(x1)
	fmt.Printf("%T\n", x1)
	fmt.Println(y1)
	fmt.Printf("%T\n", y1)
	fmt.Println(z1)
	fmt.Printf("%T\n", z1)
	fmt.Printf("================================================\n")
	//Print all variable together and call global variable

	x1, y1, z1 = 30, "Abdullah", false

	fmt.Println(x1, y1, z1)
	fmt.Printf("x = %v & Type:%T \ny = %v & Type:%T \nz = %v & Type:%T \n", x1, x1, y1, y1, z1, z1)

}
