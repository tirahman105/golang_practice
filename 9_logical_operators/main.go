package main

import "fmt"

func main() {
	var (
		a = 5
		b = 8
	)

	// &&  and
	fmt.Println(a > b && a != b)

	// ||  or
	fmt.Println(!(a >= b))

	// !   not
	fmt.Println(a == b || a < b)
}
