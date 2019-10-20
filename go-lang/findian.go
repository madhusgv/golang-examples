// String manipulation

// Main package
package main

// import go packages
import (
	"fmt"
	"strings"
)

// main function where execution starts
func main() {
	// variables
	var userInput string

	// Get the input
	fmt.Printf("Enter the findian string to match : ")
	_, err := fmt.Scan(&userInput)

	// Validate the input
	if err != nil {
		fmt.Println(err)
		return
	}
	userInputs := strings.ToLower(userInput)
	if strings.HasPrefix(userInputs, "i") &&
		strings.HasSuffix(userInputs, "n") &&
		strings.Index(userInputs, "a") != -1 {

		fmt.Println("Found!")

	} else {
		fmt.Println("Not Found!")
	}

}
