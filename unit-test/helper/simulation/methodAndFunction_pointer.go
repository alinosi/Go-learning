/**
	this code is the simulation of how unit test works in golang
**/

package main

import "fmt"

// this is the struct that will be used as a parameter for the function
type T struct {
	name    string
	number  int
	message string
}

// this is the unit test
func UnitTest(t *T) {
	if t.name != "Noby" {
		// Error handling simulation
		t.Failure("the test is failure") // calling a method if the error occurs
	}
	fmt.Println("Test success")
}

// this is the method that will change the value of the field in struct T
func (t *T) Failure(message string) {
	t.message = message
	panic(message)
}

// simulation start
func main() {

	// create an instance of struct T
	t := &T{name: "oby", number: 1}

	// run the unit test
	UnitTest(t)
}
