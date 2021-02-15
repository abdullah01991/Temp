package main

import "fmt"

/*
Following table shows all the arithmetic operators supported by Go language.
Assume variable A holds 10 and variable B holds 20
*/
// global variable x & y

var x = 50
var y = 25

func main() {
	// Local variable A & B
	A := 10
	B := 20

	fmt.Println(A + B) // 30

	fmt.Println(x + y) // 75
}
