package main

import (
	"fmt"
)

var a int

var b string = "Abdullah Almamun"

var z = 100

var y = 150

/* if we use Variable ( var z = 100 / var y = 150) outside the function
then we can call this variable value from anywhere in this calss.
*/

func main() {
	// Short declaration operator
	/* If we declar variavle inside the func body then we can use short declaration operation (x := 20)
	And if we use outside of func body the we have to use (var x = 20)
	*/

	x := 20

	fmt.Println("Print x:", x)

	var y = 30

	fmt.Println("Print y1:", y)

	fmt.Println("Print z1:", z)

	call1()
	call2()
	call3()
	//call4()

}

func call1() {
	fmt.Println("Print z2:", z)
}

func call2() {
	fmt.Println("Print y2:", y)
}

func call3() {
	fmt.Println("Print value for a:", a)

	call4()

}

func call4() {
	fmt.Println("Print value for b:", b)

}
