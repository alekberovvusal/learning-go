package main

import (
	"fmt"
)

func main() {
	var carList [3]cars

	carList[0] = cars{alarmSystem: true, make: "Opel", model: "Astra", city: "Baku", year: 2010, mileage: 30000, price: 15000, fuelType: "Diesel", color: "Black", gear: manual, status: available}
	carList[1] = cars{alarmSystem: true, make: "KIA", model: "Optima", city: "Baku", mileage: 20000, year: 2016, price: 35000, fuelType: "Gas", color: "White", gear: automatic, status: available}
	carList[2] = cars{alarmSystem: true, make: "BMW", model: "i7", city: "Baku", mileage: 10000, year: 2019, price: 70000, fuelType: "Hybrid", color: "Red", gear: automatic, status: available}

	fmt.Println("Starting..")

	fmt.Println("inital condition of the car: \n", carList[1])
	checkStatus(&carList[1])
	sellCar(&carList[1])
	checkStatus(&carList[1])
	fmt.Println("Final condition of the car: \n", carList[1])
}

type cars struct {
	// Equipments
	alarmSystem bool

	// Conditions
	status        availability
	make          string
	model         string
	color         string
	mileage       uint32
	price         uint32
	priceCurrency string
	city          string
	gear          gear
	year          uint16
	fuelType      string
}

type gear string

const (
	manual    gear = "Manual"
	automatic gear = "Automatic"
)

type availability string

const (
	available availability = "Available"
	sold      availability = "Sold"
)

func sellCar(c *cars) {
	fmt.Println("Executing sellCar function ...")
	if c.status != "Sold" {
		c.status = "Sold"
		return
	}
	fmt.Println("Car already sold. Exiting...")

}

func checkStatus(c *cars) {
	if c.status == "Sold" {
		fmt.Println("Current Status of the car: Sold!")
		return
	}
	fmt.Println("Current Status of the car: Avaliable!")
}
