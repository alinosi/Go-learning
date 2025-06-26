package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}
	values := []int{100, 95, 80, 90}

	fmt.Println(slices.Min(names)) // Mencari nilai terendah dari slices names (nilai = akumulasi byte perhuruf)
	fmt.Println(slices.Min(values)) // Mencari nilai terendah dari slices values
	fmt.Println(slices.Max(names)) // Mencari nilai tertinggi dari slices names
	fmt.Println(slices.Max(values)) // Mencari nilai tertinggi dari slices values
	fmt.Println(slices.Contains(names, "Noby")) // mencari nama noby di slice names
	fmt.Println(slices.Index(names, "Noby")) // mencari lokasi index Noby di slice names
	fmt.Println(slices.Index(names, "Paul")) // mencari lokasi index paul di slice names
}
