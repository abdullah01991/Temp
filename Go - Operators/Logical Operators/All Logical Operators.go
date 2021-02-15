package main

import "fmt"

func main() {
	var A bool = true
	var B bool = false

	if A && B {
		fmt.Printf("A && B Condition is true\n")
	} else {
		fmt.Printf("A && B Condition is false\n")
	}

	if A || B {
		fmt.Printf("A || B Condition is true\n")
	} else {
		fmt.Printf("A || B Condition is false\n")
	}

	if !(A && B) {
		fmt.Printf("!(A && B) Condition is true\n")
	} else {
		fmt.Printf("!(A && B) Condition is false\n")
	}

	fmt.Printf("==================================\n")

	var a bool = true
	var b bool = false

	if a && b {
		fmt.Printf("Line 1 - Condition is true\n")
	}
	if a || b {
		fmt.Printf("Line 2 - Condition is true\n")
	}

	/* lets change the value of  a and b */
	a = false
	b = true
	if a && b {
		fmt.Printf("Line 3 - Condition is true\n")
	} else {
		fmt.Printf("Line 3 - Condition is not true\n")
	}
	if !(a && b) {
		fmt.Printf("Line 4 - Condition is true\n")
	}
}
