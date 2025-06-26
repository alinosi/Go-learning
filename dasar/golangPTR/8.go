package main

import "fmt"

func main() {
	var a int = 42
    var p *int = &a
    fmt.Println("Nilai p:", p)
    
    if p != nil {
        fmt.Println("Nilai yang ditunjuk:", *p)
    } else {
        fmt.Println("Pointer nil!")
    }
    
    // Buat p menunjuk ke nilai 42
    // Kemudian print nilai yang ditunjuk
}