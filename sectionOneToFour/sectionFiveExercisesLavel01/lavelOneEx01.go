package main

import "fmt"

var x = 20
var y = "Abdullah Almamun"
var z = true

func main() {
	//Print all variable one by one	 and call global variable
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	//Print all variable together and call global variable
	fmt.Println(x, y, z)
	fmt.Printf("x = %v & Type:%T \ny = %v & Type:%T \nz = %v & Type:%T \n", x, x, y, y, z, z)

	//Print all variable one by one	 and call local variable
	x = 10
	y = "James Bond"
	z = true

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	//Print all variable together and call global variable

	x, y, z = 30, "Abdullah", false

	fmt.Println(x, y, z)
	fmt.Printf("x = %v & Type:%T \ny = %v & Type:%T \nz = %v & Type:%T \n", x, x, y, y, z, z)
}
