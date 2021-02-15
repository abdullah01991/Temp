package main

import "fmt"

func main() {
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}
	if 2 == 2 {
		fmt.Println("005")
	}
	if 2 != 2 {
		fmt.Println("006")
	}
	if !(2 == 2) {
		fmt.Println("007")
	}
	if !(2 != 2) {
		fmt.Println("008")
	}
	fmt.Printf("=================================================\n")
	main11()
}

func main11() {
	if x := 42; x == 2 {
		fmt.Println("This is ture x = 42")
	}
	fmt.Println("This is False x != 42")
	fmt.Println("Result something else")
	fmt.Printf("=================================================\n")

	main12()
}

func main12() {
	if x := 40; x == 40 {
		fmt.Println("This is true x = 40")
	} else {
		fmt.Println("This is false x != 40")
		fmt.Println("Result something else")
	}
}
