package main

import (
	"fmt"
	"math"
)

// GenDisplaceFn

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (((a / 2) * math.Pow(t, 2)) + (vo * t) + so)
	}
	return fn

}

// Main function
func main() {
	var a, vo, so, t float64

	// Get acceleration, Initial velocity, initial displacement

	// Get the acceleration
	fmt.Println("Please enter the acceleration( meter per second)")
	fmt.Scanln(&a)

	// Get the initial velocity
	fmt.Println("Please enter the initial velocity ( meter per second)")
	fmt.Scanln(&vo)

	// Get the initial displacement
	fmt.Println("Please enter the initial displacement ( meters )")
	fmt.Scanln(&vo)

	// Get the function to calculte displacement
	dispfn := GenDisplaceFn(a, vo, so)

	// Get the time
	fmt.Println("Please enter the time ( seconds) ")
	fmt.Scanln(&t)

	// Call the function returned by GenDisplaceFn to calculte displacement
	fmt.Println("The calculted displacement after ", t, " seconds: ", dispfn(t), "meters")
}
