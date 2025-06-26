package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.40))  // membulatkan ke atas
	fmt.Println(math.Floor(1.60)) // membulatkan ke bawah
	fmt.Println(math.Round(1.60)) // membulatkan ke bilangan terdekat
	fmt.Println(math.Max(31, 50)) // memilih angka terbesar
	fmt.Println(math.Min(31, 50)) // memilih angka terkecil
}
