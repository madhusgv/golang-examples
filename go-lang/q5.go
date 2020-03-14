package main

import (
		"fmt"
		"strings"
		"strconv"
		)

func main() {
  var xtemp int
  x1 := 0
  x2 := 1
  for x:=0; x<5; x++ {
    xtemp = x2
    x2 = x2 + x1
    x1 = xtemp
  }
  fmt.Println(x2)
  
    s := strings.Replace("ianianian", "ni", "in", 2)
  fmt.Println("String",s)
  
    i, _ := strconv.Atoi("10")
  y := i * 2
  fmt.Println("print value ",y)
  
}