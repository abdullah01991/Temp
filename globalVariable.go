package main

import "fmt"

var x int = 20

func main() {

	var x int = 100
	var y int = 15
	var z int = 0

	fmt.Printf("Value of x in Func Main : %d \n", x)
	z = sum(x, y)
	fmt.Printf("Value of z in main: %d \n", z)

}

func sum(x, y int) int {
	fmt.Printf("Value of x in Sum: %d \n", x)
	fmt.Printf("Value of y in sum: %d \n", y)
	return x + y
}
