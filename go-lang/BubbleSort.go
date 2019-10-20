package main

import (
	"fmt"
)

//swap the number to sort
func swap(data []int, idx int) {
	if data[idx-1] > data[idx] {
		temp := data[idx]
		data[idx] = data[idx-1]
		data[idx-1] = temp
	}
}

// Sort the slice
func BubbleSort(data []int) {
	for i := len(data); i > 0; i-- {
		//The inner loop will first iterate through the full length
		//the next iteration will be through n-1
		// the next will be through n-2 and so on
		for j := 1; j < i; j++ {
			swap(data, j)
		}
	}
}

// Main function
func main() {
	var n, i int

	fmt.Println("Please enter the number of inputs")
	fmt.Scanln(&n)
	if n > 10 || n <= 0 {
		fmt.Println("*** Invalid Input. Min : 1, Max : 10")
		return
	}
	fmt.Println("Please enter the inputs to sort")
	userInput := make([]int, n)
	for i = 0; i < n; i++ {
		fmt.Scanln(&userInput[i])
	}
	fmt.Println("Unsorted List", userInput)
	BubbleSort(userInput)
	fmt.Println("Sorted List", userInput)
}
