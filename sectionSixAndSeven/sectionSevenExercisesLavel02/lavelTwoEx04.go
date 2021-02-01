package main

import "fmt"

func main() {
	//Print Decimal, Binary and Hexadecimal value for a2
	a3 := 42
	fmt.Printf("Decimal:%d\t Binary:%b\t Hex:%#x\n", a3, a3, a3)

	//Shifting a2 value by 1 and printing Decimal, Binary and Hexadecimal value for b2
	b3 := a3 << 1
	fmt.Printf("Decimal:%d\t Binary:%b\t Hex:%#x\n", b3, b3, b3)
}
