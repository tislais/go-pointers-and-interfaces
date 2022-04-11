package main

import "fmt"

type Car struct {
	name string
}

type Truck struct {
	name string
}

func main() {
	bmw := Car{name: "bmw is the best car"}
	volvo := Truck{name: "volvo is the best truck"}

	// passing the whole struct and working with a copy
	bmw.name = changeNameOfCar(bmw)
	volvo.name = changeNameOfTruck(volvo)

	fmt.Println(bmw.name)
	fmt.Println(volvo.name)
}

func changeNameOfCar(c Car) string {
	return c.name + "\n\t... changed with simple function"
}

func changeNameOfTruck(t Truck) string {
	return t.name + "\n\t... changed with simple function"
}
