package main

import (
	"encoding/base64"
	"fmt"
)

// Implementation Example(concept strengthening)

// type Mahasiswa struct {
// 	name string
// }

// func (m Mahasiswa) changeName(NewNama string) string{
// 	m.name = NewNama
// 	return NewNama
// }

func main() {

	value := "Sir Noby BItzer"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decoded))
	}
	

	// Implementation example(concept strengthening)

	// anton := Mahasiswa{"anton"}
	// andhika := anton.changeName(value)

	// fmt.Println(andhika)
}
