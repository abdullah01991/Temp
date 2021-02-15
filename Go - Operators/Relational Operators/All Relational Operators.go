package main

import "fmt"

/*
All Relational Operators

(A == B) is not true
(A != B) is true.
(A > B) is not true.
(A < B) is true.
(A >= B) is not true.
(A <= B) is true.


*/

func main() {
	var A int = 10
	var B int = 20

	//(A == B) is not true
	if A == B {
		fmt.Printf(" A is equal to B \n")
		fmt.Println("true")
	} else {
		fmt.Printf("A is not equal to B \n")
		fmt.Println("false")
	}
	fmt.Printf("==================================\n")

	//(A != B) is true.
	if A != B {
		fmt.Printf("A is not equal to B \n")
		fmt.Println("true")
	} else {
		fmt.Printf("A is equal to B \n")
		fmt.Println("false")
	}
	fmt.Printf("==================================\n")

	//(A > B) is not true.
	if A > B {
		fmt.Printf("A is getter then B \n")
		fmt.Println("false")
	} else {
		fmt.Printf("A is not getter then B \n")
		fmt.Println("true")
	}
	fmt.Printf("==================================\n")

	//(A < B) is true.
	if A < B {
		fmt.Printf("A is less then B \n")
		fmt.Println("true")
	} else {
		fmt.Printf("A is not less then B \n")
		fmt.Println("false")
	}
	fmt.Printf("==================================\n")

	//(A >= B) is not true.
	if A >= B {
		fmt.Printf("A is either greater than or equal to B \n")
		fmt.Println("false")
	} else {
		fmt.Printf("A is not either greater than or equal to B \n")
		fmt.Println("true")
	}
	fmt.Printf("==================================\n")

	//(A <= B) is true.
	if A <= B {
		fmt.Printf("A is is either less than  or equal to B \n")
		fmt.Println("true")
	} else {
		fmt.Printf("A is not is either less than  or equal to B \n")
		fmt.Println("true")
	}
	fmt.Printf("==================================\n")

	var a int = 21
	var b int = 10

	if a == b {
		fmt.Printf("Line 1 - a is equal to b\n")
	} else {
		fmt.Printf("Line 1 - a is not equal to b\n")
	}
	if a < b {
		fmt.Printf("Line 2 - a is less than b\n")
	} else {
		fmt.Printf("Line 2 - a is not less than b\n")
	}
	if a > b {
		fmt.Printf("Line 3 - a is greater than b\n")
	} else {
		fmt.Printf("Line 3 - a is not greater than b\n")
	}

	/* Lets change value of a and b */
	a = 5
	b = 20
	if a <= b {
		fmt.Printf("Line 4 - a is either less than or equal to  b\n")
	}
	if b >= a {
		fmt.Printf("Line 5 - b is either greater than  or equal to b\n")
	}
}
