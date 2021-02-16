package main

import "fmt"

func main() {
	var A uint = 60 /* 60 = 0011 1100 */
	var B uint = 13 /* 13 = 0000 1101 */
	var C uint = 0

	C = A & B /* 12 = 0000 1100 (Binary AND Operator copies a bit to the result if it exists in both operands)*/
	fmt.Printf("The value of (A & B): %d \t %b\n", C, C)

	C = A | B /* 61 = 0011 1101 (Binary OR Operator copies a bit if it exists in either operand.)*/
	fmt.Printf("The value of (A | B): %d \t %b\n", C, C)

	C = A ^ B /* 49 = 0011 0001 (Binary XOR Operator copies the bit if it is set in one operand but not both.)*/
	fmt.Printf("The value of (A ^ B): %d \t %b\n", C, C)

	C = A << 2 /*  240 = 1111 0000
	(Binary Left Shift Operator. The left operands value is moved left by
	the number of bits specified by the right operand.)*/
	fmt.Printf("The value of (A << B): %d \t %b\n", C, C)

	C = A >> 2 /*  0 = 0000 1111
	(Binary Right Shift Operator. The left operands value is moved right
	by the number of bits specified by the right operand.)*/
	fmt.Printf("The value of (A >> B): %d \t %b\n", C, C)

}
