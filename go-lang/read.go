// File read.go

// Main package
package main

// import go packages
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type person struct {
	fname string
	lname string
}

func limitLength(s string, length int) string {
	if len(s) < length {
		return strings.TrimSuffix(s, "\n")
	}

	return s[:length]
}
func readFile(filen string) (err error) {

	file, err := os.Open(filen)
	defer file.Close()

	if err != nil {
		return err
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)

	var userData []person
	var line string
	var lineno int
	lineno = 0
	for {

		line, err = reader.ReadString('\n')
		lineno++
		lineArr := strings.Split(line, " ")

		if len(lineArr) == 2 {
			fname := limitLength(lineArr[0], 20)
			lname := limitLength(lineArr[1], 20)
			userData = append(userData, person{fname, lname})
		} else {
			fmt.Println("***Invalid line ", lineno, strings.TrimSuffix(line, "\n"))
		}

		if err != nil {
			break
		}

	}
	//Print the data
	for _, data := range userData {
		fmt.Println("FirstName: " + data.fname + " LastName: " + data.lname)
	}

	if err != io.EOF {
		fmt.Printf(" *** Failed!: %v\n", err)
	}

	return
}

// main function where execution starts
func main() {

	var filename string
	userInput := bufio.NewScanner(os.Stdin)

	// Get the input
	fmt.Printf("Please enter File Name :  ")
	userInput.Scan()
	filename = userInput.Text()
	err := readFile(filename)
	if err != nil && err != io.EOF {
		fmt.Printf(" *** Failed!: %v\n", err)
	}
}
