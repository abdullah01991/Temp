package main

import "fmt"

/*
1. Printing a String value and Type of value (String)
2. Printing slice of byte of that string value.
3. Printing UTF-8 length and value using (%#U)
4. Printing all index value one by one
5. Print Index position %d, Hexadecimal %#x, and UTF-8 %#U value in one Line

*/
func main() {
	stringValue := "Abdullah Almamun"
	fmt.Println(stringValue)
	fmt.Printf("Print String Value and Type: %T ==>  %v\n", stringValue, stringValue)
	fmt.Printf("================================\n")

	byteString := []byte(stringValue)
	fmt.Println(byteString)
	fmt.Printf("Print the Data Type: %T\n %v\n", byteString, byteString)
	fmt.Printf("================================\n")

	for index := 0; index < len(stringValue); index++ {
		fmt.Printf("%#U", stringValue[index])
		fmt.Printf("\n================================\n")
	}
	for index, value := range stringValue {
		fmt.Printf("At index position: %d we have Hexadecimal: %#x And UTF-8 Value: %#U\n", index, value, stringValue[index])
		fmt.Printf("================================\n")
	}

}
