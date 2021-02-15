package main

import "fmt"

/*
Following table shows all the arithmetic operators supported by Go language.
Assume variable A holds 10 and variable B holds 20
*/
// global variable x & y

func main() {
	for i := 10; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}

	}
	fmt.Printf("============================================\n")
	main1()
}

func main1() {
	A := 10
	B := 20

	fmt.Println(A % B)
	fmt.Println(B % A)
}
