package main

import (
	"fmt"
	"github.com/gnuburn/go-basic/condi"
)

func main() {
	var booleanVar bool = true
	var intVar int = 10
	var floatVar float32 = 3.14
	var stringVar string = "Hello Go"

	fmt.Printf("Boolean: %t\r\nInteger: %d\r\nFloat: %.2f\r\nString: %s\r\n",
		booleanVar, intVar, floatVar, stringVar)
	
	stringVar = "Hello GnuBurn"
	fmt.Println("New String:", stringVar)

	fmt.Println("Type of Arithemetic Operators")

	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	condi.IfTest()

	for i := 1; i < 10; i++{
    fmt.Println("number:", i)
  }

  // i := 1
  // for {
  //   fmt.Printf("number: %d\n", i)
  //   i++
  //   if i >= 10{
  //     break
  //   }
  // }

  // i = 1
  // for i < 10 {
  //   fmt.Printf("number: %d\n", i)
  //   i++
  // }
}