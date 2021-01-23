package main

import "fmt"

type Abdullah int

var a1 Abdullah
var b1 int

func main() {
	fmt.Println(a1)
	fmt.Printf("Variable Type: %T\t%v\n", a1, a1)

	fmt.Println("================================================")

	b1 = int(a1)
	fmt.Println(b1)
	fmt.Printf("Variable Type: %T\t%v\n", b1, b1)

	fmt.Println("================================================")
	main2()
}

func main2() {
	a1 = 20
	fmt.Println(a1)
	fmt.Printf("Variable Type: %T\t%v\n", a1, a1)

	fmt.Println("================================================")

	b1 = 30
	fmt.Println(b1)
	fmt.Printf("Variable Type: %T\t%v", b1, b1)

}
