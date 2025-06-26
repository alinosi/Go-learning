package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("hello/world.go")) // get directrory name
	fmt.Println(path.Base("hello/world.go")) // get file name
	fmt.Println(path.Ext("hello/world.go")) // extension of file
	fmt.Println(path.Join("hello", "world", "main.go")) // create directory with sequenc(folder....., file)
}
