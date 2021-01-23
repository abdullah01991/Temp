package main

import "fmt"

type abdullah int

var a abdullah

func main() {
	fmt.Println(a)
	fmt.Printf("Variable Type: %T\t%v\n", a, a)
	main1()
}

func main1() {
	a = 20
	fmt.Println(a)
	fmt.Printf("Variable Type: %T\t%v", a, a)
}
