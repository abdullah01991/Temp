package main

import "fmt"

//Print value fo Operator ==, <=, >=, !=, <, >
func main() {
	a1 := (42 == 42)
	b1 := (42 <= 43)
	c1 := (42 >= 43)
	d1 := (42 != 43)
	e1 := (42 < 43)
	f1 := (42 > 43)
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1)
	fmt.Println(e1)
	fmt.Println(f1)
}
