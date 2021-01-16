package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is the file will show how Control flow works in computer science")

	main1()
	fmt.Println("Print more story")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	main2()

}

func main1() {
	fmt.Println("Print Main1 function here")

	main3()
	fmt.Println("This is for fun")
}

func main2() {
	fmt.Println("This is the exit point")
}

func main3() {
	fmt.Println("This extra I just print for fun")
}
