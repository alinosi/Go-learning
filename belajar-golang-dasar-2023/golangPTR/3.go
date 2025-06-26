package main

import "fmt"

func changeValue(num *int) {
    // Ubah nilai yang ditunjuk pointer menjadi 100
	*num = 100
}

func main() {
    num := 50
    fmt.Println("Sebelum:", num)
    changeValue(&num)
    fmt.Println("Sesudah:", num)
}