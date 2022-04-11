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

	fmt.Println(bmw.name)
	fmt.Println(volvo.name)
}
