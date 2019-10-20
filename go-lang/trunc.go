// trucation of floating point number

// Main package
package main

// import go package "fmt"
import "fmt"

// main function where execution starts
func main() {
	// variables
	var userInputf float32
	var userInputi int32

	// Get the input
	fmt.Printf("Enter the floating point number : ")
	fmt.Scan(&userInputf)

	// convert float to int and print the value
	userInputi = int32(userInputf)
	fmt.Printf("The truncated version of user input : %d\n", userInputi)
}
