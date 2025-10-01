package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) IsMarried() {
	man.Name = "Sir. " + man.Name
}

func main() {
	Noby := Man{"Noby"}
	Noby.IsMarried()

	fmt.Println(Noby.Name)
}
