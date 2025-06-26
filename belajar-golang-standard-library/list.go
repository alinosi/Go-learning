package main

/**
	implementasi struktur data dobuly linked-list di golang
**/

import (
	"container/list"
	"fmt"
)

// membuat list dengan menggunakan type data List dari package list
func main() {
	var data *list.List = list.New()

	data.PushBack("Noby") // memasukkan data
	data.PushBack("nobygon")
	data.PushBack("bitzer")

	var head *list.Element = data.Front() // merujuk ke data pertama dari list
	fmt.Println(head.Value) // Noby

	next := head.Next() // data selanjutnya setelah data pertama
	fmt.Println(next.Value) // nobygon

	// dst...
	next = next.Next() // bitzer
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
