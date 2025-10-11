package belajar_golang_context

import (
	"fmt"
	"runtime"
	"testing"
)

func CreateCounters() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			destination <- counter
			counter++
		}
	}()
	return destination
}

func TestContextWithCancels(t *testing.T) {
	fmt.Println("total goroutines : ", runtime.NumGoroutine())

	destination := CreateCounters()
	for n := range destination {
		fmt.Println("COunter ", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("total goroutines :", runtime.NumGoroutine())
}

