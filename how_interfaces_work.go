package main

import (
	"fmt"
)

// Car describes the behavior of a vehicle.
type Car interface {
	Drive() string
	Stop() string
}

// Sedan represents a type of car.
type Sedan struct {
	Brand string
}

// Drive implementation for Sedan
func (s Sedan) Drive() string {
	return fmt.Sprintf("%s sedan is driving.", s.Brand)
}

// Stop implementation for Sedan
func (s Sedan) Stop() string {
	return fmt.Sprintf("%s sedan has stopped.", s.Brand)
}

// SUV represents another type of car.
type SUV struct {
	Brand string
}

// Drive implementation for SUV
func (s SUV) Drive() string {
	return fmt.Sprintf("%s SUV is driving.", s.Brand)
}

// Stop implementation for SUV
func (s SUV) Stop() string {
	return fmt.Sprintf("%s SUV has stopped.", s.Brand)
}

// When running this example, rename the tryMain function to main
// Then on terminal run go run how_interfaces_work.go
// tryMain is just main in disguise!
func tryMain() {
	// Creating instances of Sedan and SUV
	sedan := Sedan{Brand: "Toyota"}
	suv := SUV{Brand: "Jeep"}

	// Using the Car interface for generic function that operates on cars
	cars := []Car{sedan, suv}

	// Operating on cars using the interface methods
	for _, car := range cars {
		fmt.Println(car.Drive())
		fmt.Println(car.Stop())
		fmt.Println() // Add a line break for readability
	}

	var car Car
	car = SUV{Brand: "Ford"}
	fmt.Printf("It's a car %v \n", car)
	car = Sedan{Brand: "Tesla"}
	fmt.Printf("Still a car %v \n", car)

	giveMeACar(suv)
	giveMeACar(sedan)
}

// giveMeACar needs a car - whoever uses this function has to inject a car type
func giveMeACar(c Car) {
	fmt.Printf("I got a car %v \n", c)
}
