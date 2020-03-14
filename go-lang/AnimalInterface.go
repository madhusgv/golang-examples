// Main Package
package main

// Import required packages
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define Animal Interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Implement Animal Interface for Cow
type Cow struct {
}

func (c Cow) Eat()   { fmt.Println("Cow eats: grass") }
func (c Cow) Move()  { fmt.Println("Cow Locomtion Mothod: walk") }
func (c Cow) Speak() { fmt.Println("Cow Spoken sound: moo") }

// Implement Animal Interface for Bird
type Bird struct {
}

func (b Bird) Eat()   { fmt.Println("Bird eats: worms") }
func (b Bird) Move()  { fmt.Println("Bird  Locomtion Mothod: fly") }
func (b Bird) Speak() { fmt.Println("Bird Spoken sound: peep") }

// Implement Animal Interface for Snake
type Snake struct {
}

func (s Snake) Eat()   { fmt.Println("Snake eats: mice") }
func (b Snake) Move()  { fmt.Println("Snake  Locomtion Mothod: slither") }
func (b Snake) Speak() { fmt.Println("Snake Spoken sound: hsss") }

// Function to get the animal object requested
func getAnimalObject(input string, data map[string]Animal) Animal {
	for key, value := range data {
		if input == key {
			return value
		}
	}
	return nil
}

// Main function
func main() {
	// Init to get the user input
	userInput := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the Command or enter quit to terminate program")

	// DataStruture to store the user input data
	var objectStore map[string]Animal
	objectStore = make(map[string]Animal)

	// loop to get and process the user input
	for {
		fmt.Printf(">")
		userInput.Scan()
		input := userInput.Text()
		s := strings.Split(input, " ")
		// Validate the user input
		if len(s) != 3 {

			if len(s) >= 1 &&
				strings.ToLower(strings.TrimSpace(s[0])) == "quit" {
				// quit the program
				fmt.Println("Terminating Program")
				return
			}
			// Invalid user input - continue the loop
			fmt.Println("Invalid input")
		} else {
			// Process the input
			commandName := strings.ToLower(strings.TrimSpace(s[0]))
			arg1 := strings.ToLower(strings.TrimSpace(s[1]))
			arg2 := strings.ToLower(strings.TrimSpace(s[2]))

			switch commandName {

			case "newanimal":
				switch arg2 {
				case "cow":
					objectStore[arg1] = Cow{}
					fmt.Println("Created it!")
				case "bird":
					objectStore[arg1] = Bird{}
					fmt.Println("Created it!")
				case "snake":
					objectStore[arg1] = Snake{}
					fmt.Println("Created it!")
				default:
					fmt.Println("Invalid animal type")
				}
			case "query":
				currentObj := getAnimalObject(arg1, objectStore)
				if currentObj != nil {
					switch arg2 {

					case "eat":
						currentObj.Eat()
					case "move":
						currentObj.Move()
					case "speak":
						currentObj.Speak()
					default:
						fmt.Println("Invalid animal behaviour")
					}
				} else {
					fmt.Println("Invalid animal")
				}

			default:
				fmt.Println("Invalid command")
			}
		}

	}

}
