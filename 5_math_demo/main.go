package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2.4
	// b := 1.6

	c := math.Pow(a, 2)
	fmt.Printf("%.2f^%d = %.2f \n", a, 2, c)
}
