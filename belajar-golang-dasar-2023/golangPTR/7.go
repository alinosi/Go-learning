package main

import "fmt"

type NumberSlice []int

func appendToSlice(numbers *[]int) {
	// Tambahkan nilai 99 ke slice
	*numbers = append(*numbers, 99)

}

func main() {
	numbers := []int{1, 2, 3}
	fmt.Println("Sebelum:", numbers)
	appendToSlice(&numbers)
	fmt.Println("Sesudah:", numbers)
}
