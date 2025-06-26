package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	// konversi string menjadi bool
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(result, "is", reflect.TypeOf(result))
	}

	// konversi string menjadi integer
	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt, "is", reflect.TypeOf(resultInt))
	}

	// konversi integer menjadi string dengan bilangan berbasis 2
	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary, "is", reflect.TypeOf(binary))

	// konversi integer menjadi string dengan bilangan berbasis 10(basi standar)
	var stringInt string = strconv.Itoa(999)
	fmt.Println(stringInt, "is", reflect.TypeOf(stringInt))
}
