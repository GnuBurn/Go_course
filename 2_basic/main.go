package main

import (
	"fmt"
	// "github.com/gnuburn/go-basic/condi"
	// "errors"
)

type LoginError struct {
	Username string
	Message string
}

func (e *LoginError) Error() string{
	return fmt.Sprintf("Login error for user '%s': %s", e.Username, e.Message)
}

func login(username, password string) error{
	if username != "admin" || password != "password123"{
		return &LoginError{Username: username, Message: "invalid credentials"}
	}
	return nil
}

func main() {
	err := login("user", "pass")
	if err != nil{
		switch e := err.(type){
		case *LoginError: fmt.Println("Custome error occurred:", e)
		default: fmt.Println("Generic error occurred:", e)
		}
		return
	}

	fmt.Println("Login successful!")
	// result, err := divide(10, 0)
	// if err != nil{
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Result:", result)
	// var booleanVar bool = true
	// var intVar int = 10
	// var floatVar float32 = 3.14
	// var stringVar string = "Hello Go"

	// fmt.Printf("Boolean: %t\r\nInteger: %d\r\nFloat: %.2f\r\nString: %s\r\n",
	// 	booleanVar, intVar, floatVar, stringVar)
	
	// stringVar = "Hello GnuBurn"
	// fmt.Println("New String:", stringVar)

	// fmt.Println("Type of Arithemetic Operators")

	// a := 10
	// b := 3
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)
	// fmt.Println(a % b)

	// condi.IfTest()

	// for i := 1; i < 10; i++{
  //   fmt.Println("number:", i)
  // }

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