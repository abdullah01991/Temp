package main

import "fmt"

/*
init = initialize variable i := 0;
condition =  Some condition i <= 100;
post = post i++ (Added i++ means it's increment by 1)
*/
/*
func main() {
	for init; condition; post {
		fmt.Println()

	}
}
*/

func main() {
	//For FooLoop We have remember init; condition; post which is explain Above
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}
