package main

import "fmt"

/*
Following table shows all the arithmetic operators supported by Go language.
Assume variable A holds 10 and variable B holds 20
*/
// global variable x2 & y2

var x3 = 40
var y3 = 5

func main() {
	// local variable A & B
	A := 10
	B := 20

	fmt.Println(A / B)   //0
	fmt.Println(x3 / y3) //8
}
