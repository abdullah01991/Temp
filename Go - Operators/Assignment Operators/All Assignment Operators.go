package main

import "fmt"

/* ( =, +=, -=, *=, /=, %=, <<=, >>=, &=, ^=, |= )*/

var A int = 21
var C int

func main() {
	C = A

	fmt.Printf(" =  Operator Example, Value of c = %d\n", C)

	C += A
	fmt.Printf(" +=  Operator Example, Value of c = %d\n", C)

	C -= A
	fmt.Printf(" -=  Operator Example, Value of c = %d\n", C)

	C *= A
	fmt.Printf(" *=  Operator Example, Value of c = %d\n", C)

	C /= A
	fmt.Printf(" /=  Operator Example, Value of c = %d\n", C)

	C = 200

	C %= A
	fmt.Printf(" =  Operator Example, Value of c = %d\n", C)

	C <<= 2
	fmt.Printf(" <<=  Operator Example, Value of c = %d\n", C)

	C >>= 2
	fmt.Printf(" >>=  Operator Example, Value of c = %d\n", C)

	C &= 2
	fmt.Printf(" &=  Operator Example, Value of c = %d\n", C)

	C ^= 2
	fmt.Printf(" ^=  Operator Example, Value of c = %d\n", C)

	C |= 2
	fmt.Printf(" |=  Operator Example, Value of c = %d\n", C)
}
