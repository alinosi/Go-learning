package main

import (
	"fmt"
	"strconv"
)

func main() {

	counter := 0

	for i := 1; i < 1001; i++ { // lakuan pengecekan dari 1 sampai 1000
		a := strconv.Itoa(i)
		for _, values := range a {
			if values == '7' {
				counter++
			}
		}
	}

	fmt.Println(counter)

}
