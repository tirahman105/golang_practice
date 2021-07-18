package main

import "fmt"

func main() {
	//declare variable

	var str string
	var n, m int
	var mn float32

	//assign values

	str = "Hello World"
	n = 10
	m = 50
	mn = 2.25

	fmt.Println("value of str =", str)
	fmt.Println("value of n =", n)
	fmt.Println("value of m =", m)
	fmt.Println("value of mn =", mn)

	// declare and assign value to variables

	var city string = "Dhaka"
	var x int = 100

	fmt.Println("value of city =", city)
	fmt.Println("value of x =", x)

	//declare variable with defining its type

	country := "BD"
	val := 15

	fmt.Println("value of country =", country)
	fmt.Println("value of val =", val)

	//define multiple variable

	var (
		name  string
		email string
		age   int
	)
	name = "Dr. Tareq"
	email = "tirahman105@gmail.com"
	age = 30

	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(age)
}
