package main

import "fmt"

// Call All constant One by One
const variable1 = 10
const variable2 = 10.243
const variable3 = "Abdullah Almamun"

func main() {

	fmt.Println(variable1)
	fmt.Println(variable2)
	fmt.Println(variable3)
	fmt.Printf("Print Variable1 Data Type: (%T)  and Value: (%v)\n", variable1, variable1)
	fmt.Printf("Print Variable1 Data Type: (%T)  and Value: (%v)\n", variable2, variable2)
	fmt.Printf("Print Variable1 Data Type: (%T)  and Value: (%v)\n", variable2, variable2)
	abd()
}

// This are unType constant
const (
	variable4 = 20
	variable5 = 20.243
	variable6 = "Salsabil Bushra"
)

func abd() {

	fmt.Println(variable4)
	fmt.Println(variable5)
	fmt.Println(variable6)
	fmt.Printf("Print Variable4 Data Type: (%T)  and Value: (%v)\n", variable4, variable4)
	fmt.Printf("Print Variable5 Data Type: (%T)  and Value: (%v)\n", variable5, variable5)
	fmt.Printf("Print Variable6 Data Type: (%T)  and Value: (%v)\n", variable6, variable6)
	abd1()
}

// This are Type constant
const (
	variable7 int     = 30
	variable8 float64 = 30.243
	variable9 string  = "Fun"
)

func abd1() {

	fmt.Println(variable7)
	fmt.Println(variable8)
	fmt.Println(variable9)
	fmt.Printf("Print Variable7 Data Type: (%T)  and Value: (%v)\n", variable7, variable7)
	fmt.Printf("Print Variable8 Data Type: (%T)  and Value: (%v)\n", variable8, variable8)
	fmt.Printf("Print Variable9 Data Type: (%T)  and Value: (%v)\n", variable9, variable9)

}
