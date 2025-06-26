package main

import "fmt"

func swap(x, y *int) {
    temp := *x
	*x = *y
	*y = temp
}

func main() {
    x, y := 10, 20
    fmt.Printf("Sebelum: x=%d, y=%d\n", x, y)
    swap(&x, &y)
    fmt.Printf("Sesudah: x=%d, y=%d\n", x, y)
}