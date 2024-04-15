//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Directer interface {
	directVehicle()
}

type Motorcycle string

type Car string

type Truck string

func (m Motorcycle) directVehicle() {
	fmt.Println("Motorcycle is directed to the small lift")
}

func (c Car) directVehicle() {
	fmt.Println("Car is directed to the standard lift")
}

func (t Truck) directVehicle() {
	fmt.Println("Truck is directed to the large lift")
}

func directVehicles(vehicles []Directer) {
	for _, vehicle := range vehicles {
		vehicle.directVehicle()
	}
}

func main() {
	vehicles := []Directer{
		Motorcycle("Harley Davidson"),
		Car("Toyota Camry"),
		Truck("Ford F-150"),
	}
	directVehicles(vehicles)
}
