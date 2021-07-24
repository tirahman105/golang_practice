package main

import "fmt"

func main() {

	fmt.Print("Enter first number : ")
	var a int

	fmt.Scanln(&a)

	fmt.Print("Enter second number : ")
	var b int

	fmt.Scanln(&b)

	fmt.Println("First number a =", a)
	fmt.Println("Second number b =", b)

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
