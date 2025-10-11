package belajar_golang_context

import (
	"fmt"
	"testing"
	"time"
)

func testGoroutines() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("hello", i)
		}
	}()
	fmt.Println("function end")
}

func TestDoubleGo(t *testing.T) {
	testGoroutines()
	time.Sleep(5 * time.Second)
}
