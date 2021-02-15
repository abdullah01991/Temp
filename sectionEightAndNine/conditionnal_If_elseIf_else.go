package main

import "fmt"

func main() {
	x := 20
	if x == 21 {
		fmt.Println("This is true x = 21 ")
	} else if x == 22 {
		fmt.Println("This is ture x = 22 ")
	} else {
		fmt.Println("All are false")
	}
	fmt.Printf("====================================================\n")
	main13()
}

func main13() {
	y := 50
	if y == 50 {
		fmt.Println("This is ture y = 50")
	} else if y == 60 {
		fmt.Println("This is ture y = 60")
	} else if y == 70 {
		fmt.Println("This is ture y = 70")
	} else {
		fmt.Println("This is ture y = 100")
	}
}
