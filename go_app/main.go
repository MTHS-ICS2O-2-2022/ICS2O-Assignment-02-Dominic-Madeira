// Created by: Dominic M.
// Created on: March 2023
//
// This program finds the volume of a cube.
package main

import (
	"fmt"
	"math"
)

func main() {
	var sideLength float64
	var volume float64

	// input
	fmt.Println("This program finds the volume of a cube.")
	fmt.Println()
	fmt.Print("What is the side length of the cube?: ")
	fmt.Scanln(&sideLength)

	// process
	volume = math.Pow(sideLength, 3)
	// output
	fmt.Println()
	fmt.Println("The volume of the cube is:", volume, "cm³")

	fmt.Println("\nDone.")
}
