package main

import (
	"fmt"
	// "math"
)

func main() {

	// Sample input from keyboard

	fmt.Print("Enter your name:")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is :", name)

	// fmt.Println("Circle Area calculation")
	// fmt.Print("Enter a radius value: ")
	// var radius float64
	// fmt.Scanf("%f", &radius)

	// area := math.Pi * math.Pow(radius, 2)
	// fmt.Printf("Circle area with radius %.2f = %.2f \n", radius, area)

}
