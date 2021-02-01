package main

import "fmt"

const (
	a2     = 10
	b2 int = 20
)

func main() {
	//Assign Type and Untyped for const
	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Printf("Print Type:%T\t%v\n", a2, a2)
	fmt.Printf("Print Type:%T\t%v", b2, b2)

}
