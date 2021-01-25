package main

import "fmt"

/*

42420 Binary base 10
10^7=   10^6=   10^5=  10^4=10000  10^3=1000  10^2=100  10^1=10  10^0=1
	0 	     0      1	    4	       2	      4	    	2	     0

42 Binary base 2
2^7=128   2^6=64   2^5=32  2^4=16  2^3=8  2^2=4  2^1=2  2^0=1
	0 	     0        1	      0	      1	     0	    1	   0

** Numeric System 0,1,2,3,4,5,6,7,8,9,A(10),B(11),D(12),D(13),E(15),F(15), 16^1 =16
911 Hexadecimal Base 16
16^7=   16^6=   16^5=  16^4=  16^3=4096  16^2=256  16^1=16  16^0=1
	0 	   0        0	   0	  0	    	 3	       8		 F

*/

func main() {
	stringValue := "A"
	fmt.Println(stringValue)

	SliceOFByte := []byte(stringValue)
	fmt.Println(SliceOFByte)
	IndexPosition := SliceOFByte[0]
	fmt.Println(IndexPosition)
	fmt.Printf("Print Type of Value: %T , Binary: %b And HexchaDecimal value: %#x\n", IndexPosition, IndexPosition, IndexPosition)
}
