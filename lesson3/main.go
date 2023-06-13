package main

import "fmt"

var someName = "John"

func main() {
	// strings
	fmt.Println("Strings")
	var nameOne string = "Mario"
	var nameTwo = "Luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Peach"
	nameThree = "Bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Yoshi"
	fmt.Println(nameFour)

	fmt.Println(someName)

	// ints
	fmt.Println()
	fmt.Println("Integers")
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	fmt.Println()
	fmt.Println("Bits & Memory")
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 25

	fmt.Println(numOne, numTwo, numThree)

	// floats
	fmt.Println()
	fmt.Println("Floats")
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 4206969.69
	scoreThree := 1.5
	fmt.Println(scoreOne, scoreTwo, scoreThree)

}
