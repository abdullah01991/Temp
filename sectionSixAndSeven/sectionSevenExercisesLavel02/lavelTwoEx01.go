package main

import "fmt"

func main() {
	//Assign value for x1, y1, z1 and print value as Decimal, Binary, Hexadecimal
	x1 := 10
	y1 := 20
	z1 := 30

	fmt.Printf("===============Print Decinal Value============\n")
	fmt.Printf("%d\n", x1)
	fmt.Printf("%d\n", y1)
	fmt.Printf("%d\n", z1)
	fmt.Printf("===============Print Binary Value============\n")
	fmt.Printf("%b\n", x1)
	fmt.Printf("%b\n", y1)
	fmt.Printf("%b\n", z1)
	fmt.Printf("===============Print Hexcadecimal Value============\n")
	fmt.Printf("%x\n", x1)
	fmt.Printf("%x\n", y1)
	fmt.Printf("%x\n", z1)
	fmt.Printf("===============Print X-Hexcadecimal Value============\n")
	fmt.Printf("%#x\n", x1)
	fmt.Printf("%#x\n", y1)
	fmt.Printf("%#x\n", z1)
}
