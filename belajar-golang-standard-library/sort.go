package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Noby", 30},
		{"anton", 35},
		{"alinosi", 25},
		{"gunawan", 20},
	}

	// UserSlice dapat mengimplementasikan sort

	sort.Sort(UserSlice(users)) // konversi tipe data users ke UserSlice

	fmt.Println(users)
}

/*
 syarat untuk dapat menggunakan method dari package class
 kita harus mengikuti kontrak dari interface sort yaitu memiliki
 3 method : Len, Less, Swap
*/
