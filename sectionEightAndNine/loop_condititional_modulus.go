package main

import "fmt"

// https://www.tutorialspoint.com/go/go_operators.htm this the details about Go - Operators
func main() {
	for index := 10; index < 100; index++ {
		if index%3 == 0 {
			fmt.Println(index)
		}
	}
}
