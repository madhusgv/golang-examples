// slice

// Main package
package main

// import go packages
import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// main function where execution starts
func main() {
	// variables
	intSlice := make([]int, 3)
	i := 0
	for true {

		var userInput string
		fmt.Printf("Please enter a number to sort or x to quit: ")

		_, err := fmt.Scan(&userInput)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Check for exit criteria
		if strings.Compare(userInput, "X") == 0 ||
			strings.Compare(userInput, "x") == 0 {
			fmt.Println("Terminating Program...")
			return
		}
		// Convert string to integer
		x, err := strconv.Atoi(userInput)
		if err == nil {
			// Store in slice
			intSlice[i] = x
			i++
			// Sort
			intSlice1 := make([]int, len(intSlice))
			copy(intSlice1, intSlice)
			sort.Sort(sort.IntSlice(intSlice1))
			fmt.Println("Sorted Slice : ", intSlice1)
			if i > 2 {
				intSlice = append(intSlice, 1)
			}
		}
	}
}
