package main

import "fmt"

func main() {
	// For statement with single condition
	x := 0
	for x <= 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("End")

	main8()
}

// For statement with for clause
func main8() {
	for y := 0; y <= 5; y++ {
		fmt.Println(y)
	}
	main9()
}

func main9() {
	// Printing Single Condition using if statement
	z := 0
	for {
		if z > 9 {
			break
		}
		fmt.Println(z)
		z++
	}
}
