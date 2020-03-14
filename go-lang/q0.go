package main

import (
		"fmt"

		)
func main() {

i := 1

fmt.Print(i)

i++

defer fmt.Print(i+1)

fmt.Print(i)

}		
/*
func fA() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
func main() {
   fB := fA()
   fmt.Print(fB())
   fmt.Print(fB())
}
*/
/*func main() {	

	requestBody, err := json.Marshal(map[string]string{
			"NodeName":"nonstop.vm1",
			"image":"webprofile-nsjsp"})
	fmt.Println("print value ",bytes.NewBuffer(requestBody),err)
			
}*/