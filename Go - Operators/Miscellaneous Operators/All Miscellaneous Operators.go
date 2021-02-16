package main

import "fmt"

func main() {
	var A int = 10
	var B int64
	var C float64
	var ptr *int

	/* example of type operator */
	fmt.Printf("Type of variable A = %T \t \tValue = %v\n", A, A)
	fmt.Printf("Type of variable B = %T \t \tValue = %v\n", B, B)
	fmt.Printf("Type of variable C = %T\t Value = %v\n", C, C)

	/* example of & and * operators */
	ptr = &A

	fmt.Printf("Type of variable A = %T \t\t Value = %v\n", A, A)
	fmt.Printf("Type of variable *ptr = %T \t Value = %v\n", *ptr, *ptr)

}
