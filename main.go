package main

import (
	"fmt"
	"mylearning/arraydata"
	"mylearning/errorhandling"
	"mylearning/function"
	"mylearning/myinput"
	"mylearning/slice"
)

func main() {
	fmt.Println("Hello")
	myinput.CustomInput()

	var version = "1.0"
	fmt.Println(version)

	addition := function.Add(10, 20)
	fmt.Println("Addition of two numbers : ", addition)

	division := function.Divide(30.0, 20.75)
	// Formatted output
	fmt.Printf("Division of two numbers : %0.2f\n", division)

	errorhandling.ErrorHandler()

	arraydata.Array()

	slice.SliceCustom()

}
