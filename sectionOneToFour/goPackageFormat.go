package main

import "fmt"

var y3 = 20

// Group 1 fmt.Println and fmt.Printf
func main() {
	fmt.Println(y3)
	fmt.Printf("Print Type of value: %T\n", y3)
	fmt.Printf("Print Binary value of Y3: %b \n", y3)
	fmt.Printf("Print Hexadecimal value of Y3: %x \n", y3)
	fmt.Printf("Print X-Hexadecimal value of Y3: %#x \n", y3)

	y3 = 911
	fmt.Printf("Print X-Hexadecimal value of Y3: %#x \n", y3)
	// Print X-Hexadecimal, Hexadecimal, Binary, Type in one line
	fmt.Printf("%#x\t%x\t%b\t%T\n", y3, y3, y3, y3)
	fmt.Printf("Printing the defult value of y3: %v  \n", y3)

	// Group 2 fmt.Sprintf (String Print)
	s := fmt.Sprintf("Printin Sprintf: %#x\t%x\t%b\t%T\n", y3, y3, y3, y3)
	fmt.Println(s)

	// Group 3 fmt.Fprintf(We use this for Flie Printing)

	/*
	   \a   U+0007 alert or bell
	   \b   U+0008 backspace
	   \f   U+000C form feed
	   \n   U+000A line feed or newline
	   \r   U+000D carriage return
	   \t   U+0009 horizontal tab
	   \v   U+000b vertical tab
	   \\   U+005c backslash
	   \'   U+0027 single quote  (valid escape only within rune literals)
	   \"   U+0022 double quote  (valid escape only within string literals)
	*/
	fmt.Printf("Start From Here\n")
	fmt.Printf("%b\a%x\n", y3, y3)
	fmt.Printf("%b\b%x\n", y3, y3)
	fmt.Printf("%b\f%x\n", y3, y3)
	fmt.Printf("%b\n%x\n", y3, y3)
	fmt.Printf("%b\r%x\n", y3, y3)
	fmt.Printf("%b\t%x\n", y3, y3)
	fmt.Printf("%b\v%x\n", y3, y3)
	fmt.Printf("%b\\%x\n", y3, y3)
	fmt.Printf("%b\"%x\n", y3, y3)

}
