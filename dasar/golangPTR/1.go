package main

import "fmt"

func main() {
	var x int = 42
	// Buat pointer yang menunjuk ke variabel x
	// Kemudian print alamat memori dan nilai yang ditunjuk

	var c *int = &x
	fmt.Println(c)
}
