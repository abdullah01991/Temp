package main

import "fmt"

var x bool

func main() {
	fmt.Println(x)
	x = true

	fmt.Println(x)
	fmt.Printf("Type of x: %T\n", x)

	var y bool = true
	fmt.Println(y)

	fmt.Printf("==============================\n")

	main1()
}

func main1() {
	a := 20
	b := 20
	fmt.Printf("1. Print bool value of a & b (a == b): %v\n", a == b)

	fmt.Printf("==============================\n")

	a = 10
	b = 15

	fmt.Printf("2. Print bool value of a & b (a != b): %v\n", a != b)

	fmt.Printf("==============================\n")

	a = 10
	b = 15

	fmt.Printf("2. Print bool value of a & b (a <= b): %v\n", a <= b)

	fmt.Printf("==============================\n")

	a = 10
	b = 15

	fmt.Printf("2. Print bool value of a & b (a >= b): %v\n", a >= b)

}
