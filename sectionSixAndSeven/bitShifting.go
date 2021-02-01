package main

import (
	"fmt"
)

func main() {
	// Shifting Value base 10 and Binary value shifting value by 1
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	fmt.Printf("====================================================\n")

	main6()
}

func main6() {
	// Shifting Value base 10 and Binary value shifting value by 10
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
	fmt.Printf("====================================================\n")
	main7()
}

const (
	// Shifting Value base 10 and Binary value shifting value by 10 using Iota
	_  = iota             // iota = 0
	kb = 1 << (iota * 10) // iota = 1
	mb = 1 << (iota * 10) // iota = 2
	gb = 1 << (iota * 10) // iota = 3
)

func main7() {
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
