package main

import (
	"fmt"
	"runtime"
)

/*
int = no decimal point ( Which also called whole number)
floating number = has decimal point ( Which is called real number)
Example = main() and Example = main2()
*/

/*In Golang if we run int number (int x = 20.07428) float64 and
float64 number (float64 y = 20). It will give error

*/
/*
In Golang we use int8 (-128 to 127) for Negative to Positive both number
But Uint8 (0 to 255) we use for only positive number.
Example = main3()
*/

func main() {
	x := 20
	y := 20.234

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("Type of x %T\n", x)
	fmt.Printf("Type of y %T\n", y)
	fmt.Printf("================================\n")
	main2()
}

var x1 int
var y1 float64

func main2() {

	x1 = 10
	y1 = 10.3242

	fmt.Println(x1)
	fmt.Println(y1)
	fmt.Printf("Type of x1 %T\n", x1)
	fmt.Printf("Type of y1 %T\n", y1)
	fmt.Printf("================================\n")

	main3()
}

/*
In Golang we use int8 (-128 to 127) for Negative to Positive both number
But Uint8 (0 to 255) we use for only positive number.
Example = main3()
*/

var x2 int8 = -20
var y2 uint8 = 20

func main3() {
	fmt.Println(x2)
	fmt.Println(y2)
	fmt.Printf("Type of x2 %T\n", x2)
	fmt.Printf("Type of y2 %T\n", y2)
	fmt.Printf("================================\n")

	main4()

}

/*
In Golang we use byte is same as uint8.
And rune is same as int32.
Example = main4()
*/

var x3 byte
var y3 rune

func main4() {

	x3 = 30
	y3 = 35
	fmt.Println(x3)
	fmt.Println(y3)
	fmt.Printf("Type of x2 %T\n", x3)
	fmt.Printf("Type of y2 %T\n", y3)
	fmt.Printf("================================\n")

	main5()
}

// We can use this to check runtime system.
func main5() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOROOT())
}
