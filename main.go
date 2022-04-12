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

	// passing reference to original struct and working with that
	changeNameOfCarUsingPointer(&bmw)
	changeNameOfTruckUsingPointer(&volvo)

	bmw.name = bmw.changeNameUsingMethod()
	volvo.name = volvo.changeNameUsingMethod()

	bmw.changeNameUsingMethodWithPointer()
	volvo.changeNameUsingMethodWithPointer()

	fmt.Println(bmw.name)
	fmt.Println(volvo.name)
}

// working with a copy
func changeNameOfCar(c Car) string {
	return c.name + "\n\t... changed with simple function"
}

func changeNameOfTruck(t Truck) string {
	return t.name + "\n\t... changed with simple function"
}

// working with pointers
func changeNameOfCarUsingPointer(c *Car) {
	c.name = c.name + "\n\t... changed using pointer"
}

func changeNameOfTruckUsingPointer(t *Truck) {
	t.name = t.name + "\n\t... changed using pointer"
}

// using a method
func (c Car) changeNameUsingMethod() string {
	return c.name + "\n\t... changed with method"
}

func (t Truck) changeNameUsingMethod() string {
	return t.name + "\n\t... changed with method"
}

func (c *Car) changeNameUsingMethodWithPointer() {
	c.name = c.name + "\n\t... changed with method using pointer"
}

func (t *Truck) changeNameUsingMethodWithPointer() {
	t.name = t.name + "\n\t... changed with method using pointer"
}
