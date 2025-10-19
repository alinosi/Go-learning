package belajar_golang_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "dibatalkan:", ctx.Err())
			return
		default:
			fmt.Println(name, "sedang bekerja...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func TestNestedContext(t *testing.T) {
	parentCtx, parentCancel := context.WithCancel(context.Background())
	childCtx, childCancel := context.WithCancel(parentCtx)
	anotherChildCtx, anotherChildCancel := context.WithCancel(parentCtx)

	go worker(parentCtx, "ParentWorker")
	go worker(childCtx, "ChildWorker")
	go worker(anotherChildCtx, "AnotherWorker")

	time.Sleep(2 * time.Second)
	fmt.Println("Membatalkan child context...")
	childCancel() // hanya child yang dibatalkan

	time.Sleep(2 * time.Second)
	fmt.Println("Membatalkan parent context...")
	parentCancel() // membatalkan semuanya (termasuk turunan)

	time.Sleep(2 * time.Second)
	fmt.Println("Membatalkan anotherChild context...")
	anotherChildCancel() // membatalkan semuanya (termasuk turunan)

	time.Sleep(1 * time.Second)
}
