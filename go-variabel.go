package main

import "fmt"

func main() {
	// It's a variable declaration.
	var name string

	name = "ulil albab"
	fmt.Println(name)

	name = "albab ulil"
	fmt.Println(name)

	// It's a variable declaration.
	var friendName string = "Budi"
	fmt.Println(friendName)

	// It's a shorthand declaration.
	var age = 20
	fmt.Println(age)

	// It's a shorthand declaration.
	country := "Indonesia"
	fmt.Println(country)

	country = "United States"
	fmt.Println(country)
	// It's a multiple variable declaration.
	var (
		fristName = "John"
		lastName  = "albab"
	)
	fmt.Println(fristName, lastName)
}
