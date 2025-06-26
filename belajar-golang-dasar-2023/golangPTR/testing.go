package main

import "fmt"

func changeb(b int) {
	b = 12
}

func main() {

	var b int = 12
	changeb(b)

	fmt.Println(b) // hanya copynya

	b = 1 // b asli
	// *b = 1  ->  tidak melakukan ini untuk mengubah nilai b

	a := 12
	c := &a
	*c = 12 // ini baru benar

	// kalau mau menggunakan astrisk operator, nilanya harus berupa pointer(berawalan &)

	fmt.Println(b)
}

// pointer hanya digunakan ketika ingin merubah value dari sebuah variabel yang scopenya berbeda
// jika masih 1 scope(cth: sama-sama dalam func main) maka tidak perlu menggunakan ptr untuk merujuk ke variabel asli
