package main

import "fmt"

func main() {

	var (
		a = 5
		b = 12
	)

	fmt.Printf(" a = %d \n", a)
	fmt.Printf(" b = %d \n", b)

	fmt.Print(" Is a greater than b ? ")
	// a is greater than b

	fmt.Println(a > b)

	fmt.Print(" Is a less than b ? ")
	// a is less than b
	fmt.Println(a < b)

	fmt.Print(" Is a greater than or equal to b ? ")
	// a is greater than or equal to b
	fmt.Println(a >= b)

	fmt.Print(" Is a less than or equal to b ? ")
	// a is less than or euaal to b
	fmt.Println(a <= b)

	fmt.Print(" Is a not equal to b ? ")
	// a is not equal to b
	fmt.Println(a != b)

	fmt.Print(" Is a equal to b ? ")
	// a is equal to b
	fmt.Println(a == b)

}
