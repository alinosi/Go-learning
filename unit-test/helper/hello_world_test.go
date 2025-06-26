package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldNoby(t *testing.T) {
	result := Hello("Noby")

	if result != "Hello Noby" {
		// error
		t.Error("Result must be 'Hello Noby'")
	}

	fmt.Println("TestHelloWorldNoby Done")
}

func TestHelloWorldBitzer(t *testing.T) {
	result := Hello("Bitzer")

	if result != "Hello Bitzer" {
		// error
		t.Fatal("Result must be 'Hello Bitzer'")
	}

	fmt.Println("TestHelloWorldBitzer Done")
}
