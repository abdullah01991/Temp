package main

import "fmt"

var y2 string

var x2 int

func main() {

	fmt.Println("The value of y2:", y2)
	fmt.Printf("The Type of y2: %T \n", y2)

	y2 = "Abdullah Almamun"

	fmt.Println("The value of y2:", y2)
	fmt.Printf("The Type of y2: %T \n", y2)

	fmt.Println("The value of x2:", x2)
	fmt.Printf("The Type of x2: %T \n", x2)

	x2 = 20

	fmt.Println("The value of x2:", x2)
	fmt.Printf("The Type of x2: %T \n", x2)
}
