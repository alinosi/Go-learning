package main

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	sayHello := func() {
		fmt.Println("hello world")
	}

	sayHello()

	sayHelloName := func(name string) {
		fmt.Printf("hello world %s", name)
	}

	sayHelloName("amdhika")
}
