package main

import "fmt"

func main() {
	x := 0

	for {
		x++
		if x > 100 {
			break
		}
		if x%2 != 0 {
			continue
		}

		fmt.Println(x)

	}
	fmt.Println("End Here")

	main10()
}

func main10() {
	a := 0
	for {
		a++
		if a > 50 {
			break
		}
		if a%2 == 0 {
			continue
		}
		fmt.Println(a)
	}
}
