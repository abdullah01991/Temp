package main

import "fmt"

func main() {
	/*
		Following table shows all the arithmetic operators supported by Go language.
		Assume variable A holds 10 and variable B holds 20
	*/
	var A = 10
	var B = 20
	var C int

	C = A + B
	fmt.Printf("Adds value of C is: %v\n", C)

	C = A - B
	fmt.Printf("Subtracts value of C is: %v\n", C)

	C = A * B
	fmt.Printf("Multiplies value of C is: %v\n", C)

	C = A / B
	fmt.Printf("Divides value of C is: %v\n", C)

	C = A % B
	fmt.Printf("Modulus value of C is: %v\n", C)

	A++
	fmt.Printf("Increment value of A is: %v\n", A)

	B++
	fmt.Printf("Increment value of A is: %v\n", B)

	A--
	fmt.Printf("Decrement value of A is: %v\n", A)

	B--
	fmt.Printf("Decrement value of A is: %v\n", B)

	fmt.Printf("=========================================\n")
	var a int = 10
	var b int = 20
	var c int

	c = a + b
	fmt.Printf("Line 1 - Value of c is %d\n", c)

	c = a - b
	fmt.Printf("Line 2 - Value of c is %d\n", c)

	c = a * b
	fmt.Printf("Line 3 - Value of c is %d\n", c)

	c = a / b
	fmt.Printf("Line 4 - Value of c is %d\n", c)

	c = a % b
	fmt.Printf("Line 5 - Value of c is %d\n", c)

	a++
	fmt.Printf("Line 6 - Value of a is %d\n", a)

	a--
	fmt.Printf("Line 7 - Value of a is %d\n", a)
}
