package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func updateAge(p *Person) {
    p.Age++ // it's equivalent to (*p).Age
}

func main() {
    p := Person{Name: "Alice", Age: 25}
    fmt.Printf("Sebelum: %+v\n", p)
    updateAge(&p)
    fmt.Printf("Sesudah: %+v\n", p)
}