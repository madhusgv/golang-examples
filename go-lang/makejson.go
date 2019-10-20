// slice

// Main package
package main

// import go packages
import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// main function where execution starts
func main() {
	userData := make(map[string]string)
	var name string
	var addr string
	userInput := bufio.NewScanner(os.Stdin)

	// Get the input
	fmt.Printf("Enter the name :  ")
	userInput.Scan()
	name = userInput.Text()

	fmt.Printf("Enter the Address :  ")
	userInput.Scan()
	addr= userInput.Text()

	userData["name"] = name
	userData["address"] = addr
	barr, err := json.Marshal(userData)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User Information in Map :  ", userData)
	fmt.Println("User Information in JSON :  ", string(barr))

}
