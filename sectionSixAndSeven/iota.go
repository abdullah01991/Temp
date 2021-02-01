package main

import "fmt"

/*
Iota is incrementing value by 1
In Golan if we need to increment something by one We can use (iota)
*/

const (
	variable10 = iota
	variable11 = iota
	variable12 = iota
)

func main() {

	fmt.Println(variable10)
	fmt.Println(variable11)
	fmt.Println(variable12)
	fmt.Printf("Print Variable10 Data Type: (%T)  and Value: (%v)\n", variable10, variable10)
	fmt.Printf("Print Variable11 Data Type: (%T)  and Value: (%v)\n", variable11, variable11)
	fmt.Printf("Print Variable12 Data Type: (%T)  and Value: (%v)\n", variable12, variable12)
	abd2()
}

// Iota will print same value like in the above
const (
	variable13 = iota
	variable14
	variable15
)

func abd2() {

	fmt.Println(variable13)
	fmt.Println(variable14)
	fmt.Println(variable15)
	fmt.Printf("Print variable13 Data Type: (%T)  and Value: (%v)\n", variable13, variable13)
	fmt.Printf("Print Variable14 Data Type: (%T)  and Value: (%v)\n", variable14, variable14)
	fmt.Printf("Print Variable15 Data Type: (%T)  and Value: (%v)\n", variable15, variable15)
}
