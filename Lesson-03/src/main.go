package main

import (
	"fmt"
)

func main() {
	var carList [11]cars

	carList[0] = cars{alarmSystem: true, make: "Opel", model: "Astra", city: "Baku", year: 2010, mileage: 30000, price: 15000, fuelType: "Diesel", color: "Black", gear: manual, status: available}
	carList[1] = cars{alarmSystem: true, make: "KIA", model: "Optima", city: "Baku", mileage: 20000, year: 2016, price: 35000, fuelType: "Gas", color: "White", gear: automatic, status: available}
	carList[2] = cars{alarmSystem: true, make: "BMW", model: "i7", city: "Baku", mileage: 10000, year: 2019, price: 70000, fuelType: "Hybrid", color: "Red", gear: automatic, status: available}
	carList[3] = cars{alarmSystem: true, make: "Subaru", model: "Vectra", city: "Baku", year: 2010, mileage: 30000, price: 15000, fuelType: "Diesel", color: "Black", gear: manual, status: available}
	carList[4] = cars{alarmSystem: true, make: "Toyota", model: "Rio", city: "Baku", mileage: 20000, year: 2016, price: 35000, fuelType: "Gas", color: "White", gear: automatic, status: available}
	carList[5] = cars{alarmSystem: true, make: "Mercedes", model: "x5", city: "Baku", mileage: 10000, year: 2019, price: 70000, fuelType: "Hybrid", color: "Red", gear: automatic, status: available}
	carList[6] = cars{alarmSystem: true, make: "Lada", model: "2107", city: "Baku", year: 2010, mileage: 30000, price: 15000, fuelType: "Diesel", color: "Black", gear: manual, status: available}
	carList[7] = cars{alarmSystem: true, make: "Jeep", model: "Cerato", city: "Baku", mileage: 20000, year: 2016, price: 35000, fuelType: "Gas", color: "White", gear: automatic, status: available}
	carList[8] = cars{alarmSystem: true, make: "Tesla", model: "525", city: "Baku", mileage: 10000, year: 2019, price: 70000, fuelType: "Hybrid", color: "Red", gear: automatic, status: available}
	carList[9] = cars{alarmSystem: true, make: "Ford", model: "Astra", city: "Baku", year: 2010, mileage: 30000, price: 15000, fuelType: "Diesel", color: "Black", gear: manual, status: available}
	carList[10] = cars{alarmSystem: true, make: "Cadillac", model: "Koup", city: "Baku", mileage: 20000, year: 2016, price: 35000, fuelType: "Gas", color: "White", gear: automatic, status: available}

	fmt.Println("Starting..")

	//fmt.Println("inital condition of the car: \n", carList)
	// checkStatus(&carList[1])
	// sellCar(&carList[1])
	// checkStatus(&carList[1])
	// fmt.Println("Final condition of the car: \n", carList[1])

	searchCar(2015, carList[:])
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

func searchCar(y uint16, a []cars) {
	var k int
	var matchedCars []string

	for k = 0; k < len(a); k++ {
		if a[k].year >= y {
			matchedCars = append(matchedCars, a[k].make)
			//	fmt.Printf("Brand: %v Model: %v Year: %v \n", a[k].make, a[k].model, a[k].year)

		}

	}
	fmt.Println(matchedCars)
}
