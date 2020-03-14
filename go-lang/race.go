package main

import (
	"fmt"
	"sync"
)

var Wait sync.WaitGroup
var Counter int = 0

func main() {

	for id := 1; id <= 2; id++ {

		Wait.Add(1)
		go Routine(id)
	}

	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}

func Routine(id int) {

	for count := 0; count < 2; count++ {

		value := Counter
		value++
		Counter = value
	}

	Wait.Done()
}
