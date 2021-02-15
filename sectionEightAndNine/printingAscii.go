package main

import "fmt"

// Here I am printing Decimal, Binary, HexDecimal, Ascii format from 33 to 122.
func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("Decimal:- %d\t Binary:- %b\t HexDecimal:- %#x\t Ascii:- %#U\n", i, i, i, i)
	}
}
