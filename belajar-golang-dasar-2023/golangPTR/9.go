package main

import (
	"fmt"
)

func changePointer(ptr **int) {
	// Buat pointer menunjuk ke nilai baru (88)
	newValue := 88
	*ptr = &newValue
	// Ubah pointer agar menunjuk ke newValue
}

func main() {
	x := 42
	ptr := &x

	fmt.Println("Sebelum:", *ptr)
	changePointer(&ptr)
	fmt.Println("Sesudah:", *ptr)
}

// untuk soal ini apakah walau parameter yang diterima adalah pointer-pointer.
// compiler akan langsung mengenali walau kita hanya menggunakan pointer
// walau di lingkungan fungsinya yang didefinisikan adlaah pointer pointer?
