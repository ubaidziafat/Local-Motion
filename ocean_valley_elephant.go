package main

import (
	"fmt"

)

func main() {

	// Print Information about Local Motion //
	fmt.Println("Local Motion is a company that develops software for social mobility, electric bike-sharing, and e-scooter sharing services.")
	
	// Create and Initialize Variables //
	var bikes int
	bikes = 100
	var scooters int
	scooters = 50
	var vehicleCode string
	vehicleCode = "VN10"
	
	// Write a Function to Calculate the Total Number of Vehicles //
	func totalVehicles() int{
		total := bikes + scooters
		return total
	}
	
	// Print out the Total Number of Vehicles and Vehicle Code. 
	fmt.Println("Local Motion has a total of", totalVehicles(), " vehicles with the vehicle code of", vehicleCode, ".")
	
	// Create a New Struct to Store Vehicle Information //
	type Vehicle struct{
		name string
		quantity int
		code string
		location string
		
	}
	
	// Create Instances of the Struct //
	bike1 := Vehicle {
		name: "Bike",
		quantity: bikes,
		code: vehicleCode,
		location: "Berlin",
	}
	
	scooter1 := Vehicle {
		name: "Scooter",
		quantity: scooters,
		code: vehicleCode,
		location: "Berlin",
	}
	
	// Print Out Details About the Different Instances //
	fmt.Println("Details about Bike Instance:")
	fmt.Println("Name:", bike1.name)
	fmt.Println("Quantity:", bike1.quantity)
	fmt.Println("Vehicle Code:", bike1.code)
	fmt.Println("Location:", bike1.location)
	
	fmt.Println("Details about Scooter Instance:")
	fmt.Println("Name:", scooter1.name)
	fmt.Println("Quantity:", scooter1.quantity)
	fmt.Println("Vehicle Code:", scooter1.code)
	fmt.Println("Location:", scooter1.location)
	
	// Write a Function to Increase the Quantity of a Vehicle //
	func increaseQuantity(v Vehicle, q int) Vehicle {
		v.quantity += q
		return v
	}
	
	// Increase the Quantity of Bikes by 10 //
	bike1 = increaseQuantity(bike1, 10)
	
	// Increase the Quantity of Scooters by 20 //
	scooter1 = increaseQuantity(scooter1, 20)
	
	// Print Out the New Quantities of the Vehicle Instances //
	fmt.Println("After Increase in Quantities:")
	fmt.Println("Bikes:", bike1.quantity)
	fmt.Println("Scooters:", scooter1.quantity)
	
	// Write a Method to Change the Location of a Vehicle //
	func (v Vehicle) changeLocation(location string) {
		v.location = location
	}
	
	// Change the Location of the Bike Instance to London //
	bike1.changeLocation("London")
	
	// Change the Location of the Scooter Instance to Paris //
	scooter1.changeLocation("Paris")
	
	// Print Out the New Locations of the Instances //
	fmt.Println("After Change of Locations:")
	fmt.Println("Bike Location:", bike1.location)
	fmt.Println("Scooter Location:", scooter1.location)
	
	// Create a Slice of Vehicles //
	vehicles := []Vehicle{bike1, scooter1}
	
	// Write a Function to Calculate the Total Number of Vehicles Across All Locations //
	func getTotalNumberAcrossLocations(v []Vehicle) int{
		var total int
		for _, vehicle := range v {
			total += vehicle.quantity
		}
		return total
	}
	
	// Print Out the Total Number of Vehicles Across All Locations //
	fmt.Println("Local Motion has a total of", getTotalNumberAcrossLocations(vehicles), "vehicles across all locations.")
	
}