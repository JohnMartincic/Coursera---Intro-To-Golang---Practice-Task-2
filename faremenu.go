package main

import "fmt"

func main() {
	var origin, destination, cabinClass string
	const enteredMsg = "You've entered "

	fmt.Println("*** Silver Karma Airfare Calculator ***")

	validOriginEntered := false
	var originCity City
	var originErr error

	for !validOriginEntered {
		fmt.Println("Enter origin code: ")
		fmt.Scanln(&origin)

		originCity, originErr = getCityFromCode(origin)

		if originErr == nil {
			fmt.Println(enteredMsg + originCity.cityName)
			validOriginEntered = true
		} else {
			fmt.Println(originErr)
		}
	}

	validDestinationEntered := false
	var destinationCity City
	var destinationErr error

	for !validDestinationEntered {
		fmt.Println("Enter destination code: ")
		fmt.Scanln(&destination)

		destinationCity, destinationErr = getCityFromCode(destination)

		if destinationErr == nil {
			fmt.Println(enteredMsg + destinationCity.cityName)
			validDestinationEntered = true
		} else {
			fmt.Println(destinationErr)
		}
	}

	validCabinClassEntered := false
	var enteredCabinClass CabinClass
	var enteredCabinClassErr error

	for !validCabinClassEntered {
		for i := range cabinClasses {
			fmt.Println(cabinClasses[i].code + "=" + cabinClasses[i].className)
		}

		fmt.Println("Enter cabin class code: ")
		fmt.Scanln(&cabinClass)

		enteredCabinClass, enteredCabinClassErr = getCabinClassFromCode(cabinClass)

		if enteredCabinClassErr == nil {
			fmt.Println(enteredMsg + enteredCabinClass.className + " class")
			validCabinClassEntered = true
		} else {
			fmt.Println(enteredCabinClassErr)
		}
	}

	distance := CalculateDistance(float64(destinationCity.longitude)/10000,
		float64(destinationCity.latitude)/10000,
		float64(originCity.longitude)/10000,
		float64(originCity.latitude)/10000)

	fmt.Printf("\nDtance = %.1f km\n", distance)
	rateFloat32 := float32(enteredCabinClass.rate) / 100

	fmt.Printf("$ per km = %.2f\n", rateFloat32)
	fmt.Printf("Total fare = $%.2f\n", rateFloat32*float32(distance))
}
