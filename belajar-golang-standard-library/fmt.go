package main

import "fmt"

func main() {
	firstName := "Noby"
	lastName := "nobygon"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}

// fmt package contains many formats 
// that can be used to make writing more structured.