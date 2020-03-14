package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() string {
	return a.food
}
func (a Animal) Move() string {
	return a.locomotion
}
func (a Animal) Speak() string {
	return a.noise
}

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}
	userInput := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter input to display animal behaviours or enter quit to terminate program")

	for {
		fmt.Printf(">")
		userInput.Scan()
		input := userInput.Text()
		s := strings.Split(input, " ")
		if strings.Compare(s[0], "cow") == 0 {
			if strings.Compare(s[1], "eat") == 0 {
				fmt.Println(s[0] + " " + s[1] + ": " + cow.Eat())
			} else if strings.Compare(s[1], "move") == 0 {
				fmt.Println(s[0] + " " + s[1] + ": " + cow.Move())
			} else if strings.Compare(s[1], "speak") == 0 {
				fmt.Println(s[0] + " " + s[1] + ": " + cow.Speak())
			} else {
				fmt.Println("***Error Invalid Input")
			}

		} else if strings.Compare(s[0], "bird") == 0 {
			if strings.Compare(s[1], "eat") == 0 {
				fmt.Println(s[0] + " " + s[1] + ": " + bird.Eat())
			} else if strings.Compare(s[1], "move") == 0 {
				fmt.Println(s[0] + " " + s[1] + ": " + bird.Move())
			} else if strings.Compare(s[1], "speak") == 0 {
				fmt.Println(s[0] + " " + s[1] + ": " + bird.Speak())
			} else {
				fmt.Println("***Error Invalid Input")
			}
		} else if strings.Compare(s[0], "snake") == 0 {
			if strings.Compare(s[1], "eat") == 0 {
				fmt.Println(s[0] + " " + s[1] + ": " + snake.Eat())
			} else if strings.Compare(s[1], "move") == 0 {
				fmt.Println(s[0] + " " + s[1] + ": " + snake.Move())
			} else if strings.Compare(s[1], "speak") == 0 {
				fmt.Println(s[0] + " " + s[1] + ": " + snake.Speak())
			} else {
				fmt.Println("***Error Invalid Input")
			}
		} else if strings.Compare(input, "quit") == 0 {
			return
		} else {
			fmt.Println("***Error Invalid input")
		}

	}

}
