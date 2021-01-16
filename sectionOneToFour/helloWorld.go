package main

import "fmt"

func main() {
	var b = 0
	n, err := fmt.Println("Hello World", 40, true)
	fmt.Println(n)
	fmt.Println(err)
	fmt.Println(b)
}
