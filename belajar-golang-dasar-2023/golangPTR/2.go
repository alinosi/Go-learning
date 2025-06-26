package main

import "fmt"

func main() {
	a := 10
	b := &a
	c := &b

	fmt.Println(a)  // outputnya adalah 10 karena deklarasi value langsung
	fmt.Println(*b) // outputnya adalah 10 karena asterisk operator mengkonversi pointer menjadi nilai aktual.
	fmt.Println(**c)
	// outputnya adalah 10 karena c menyimpan pointer dari b dan b menyimpan pointer dari a,
	// dengan double asterisk operator maka alurnya menjadi konversi pointer b menjadi b
	// yang mana nilai b adalah pointer a lalu konversi pointer a menjadi a yang mana nilainya adalah 10

}
