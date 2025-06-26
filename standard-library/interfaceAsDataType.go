package main

import (
	"fmt"
	"reflect"
)

type Vehicle interface {
	start()
}

type Car struct {
	name string
}

type Bike struct {
	name string
}

func (c Car) start() {
	fmt.Println("open the car")
	fmt.Println("sit on the seat")
	fmt.Println("Car is starting")
}

func (c Bike) start() {
	fmt.Println("sit on the seat")
	fmt.Println("Bike is starting")
}

func brandCheck(value Vehicle) {
	value.start()
}

func returnVehicle() Vehicle {
	return Car{name: "Toyota"}
}

func main() {
	var a Vehicle = Car{name: "volvo"} // initialization variable with interface data type
	b := Car{"mercedes"}               // normal intialization, struct as data type
	var c Vehicle = returnVehicle()    // Interface as data type

	// In this case, this is the example of wrong way to determine variable data type
	// var d Car = returnVehicle()

	// it's occur because a vehicle is not necessarily a car.
	// It could be that the returned data type is a bike not a car.


	fmt.Println("VAR a")
	a.start()
	brandCheck(a)                  // volvo
	fmt.Println(a)                 // {volvo}
	fmt.Println(reflect.TypeOf(a)) // ouput : Car
	fmt.Println()

	fmt.Println("VAR b")
	b.start()
	brandCheck(b)                  // mercedes
	fmt.Println(b)                 // {mercedes}
	fmt.Println(reflect.TypeOf(b)) // ouput : Car
	fmt.Println()

	fmt.Println("VAR c")
	c.start()
	brandCheck(c)                  // Toyota
	fmt.Println(c)                 // {Toyota}
	fmt.Println(reflect.TypeOf(c)) // ouput : Car
}

// data type of a is Car even though a is declared as Vehicle, compiler recognizes a as Car

// data type of b is Car because that was direct declaration

// data type of c is Car even though c is declared as Vehicle, compiler recognizes c as Car

// but the start() method has different references based on how the caller initialized it

// a is initialized with an Vehicle(interface) as the data type. start() owned by interface
// b is initialized with an Car(struct) as the data type. start() owned by Car
// a is initialized with an Vehicle(interface) as the data type. start() owned by interface
