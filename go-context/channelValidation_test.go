package belajar_golang_context

import (
	"fmt"
	"runtime"
	"testing"
)

func CreateCounterValidation() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			destination <- counter
			counter++
		}
	}()
	return destination // return channel that are still active
}

func TestContextWithCancelValidation(t *testing.T) {
	fmt.Println("total goroutines : ", runtime.NumGoroutine())

	destination := CreateCounterValidation()
	for n := range destination {
		fmt.Println("Counter ", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("total goroutines :", runtime.NumGoroutine())
}

// number of counters will be equal to
