package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`n([a-z])([a-z])y`) // n-[a-z]-[a-z]-y (letter combination) 

	fmt.Println(regex.MatchString("noby")) // true
	fmt.Println(regex.MatchString("novy")) // true
	fmt.Println(regex.MatchString("noby")) // true 
	fmt.Println(regex.MatchString("nobams")) // false

	fmt.Println(regex.FindAllString("noby novy n0bu nony noBy", 10))
}
