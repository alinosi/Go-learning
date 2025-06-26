package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("hello/world.go"))
	fmt.Println(filepath.Base("hello/world.go"))
	fmt.Println(filepath.Ext("hello/world.go"))
	fmt.Println(filepath.IsAbs("hello/world.go"))   // is absolute, Ex: C:\Users\User> (start from the top directory)
	fmt.Println(filepath.IsLocal("hello/world.go")) // is relative
	fmt.Println(filepath.Join("hello", "world", "main.go"))
}
