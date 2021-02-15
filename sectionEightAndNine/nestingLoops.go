package main

import "fmt"

func main() {
	//Printing Inner and outer loop
	for i := 0; i <= 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Printing Outer Loop: %d\t Printing Inner Loop: %d\t\n", i, j)
		}
		fmt.Printf("====================================================\n")
	}
	main7()
}

func main7() {
	for k := 0; k <= 6; k++ {
		fmt.Printf("Printion Outer Loop: %d\n", k)
		for l := 0; l < 2; l++ {
			fmt.Printf("\t\tPrintion Inner Loop: %d\n", l)
		}
		fmt.Printf("====================================================\n")
	}
}
