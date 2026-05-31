package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gnuburn/go-syntax/gnu"
)

func sayHello(){
	fmt.Println("Hello World")
}

func main(){
	sayHello()
	id := uuid.New()
	fmt.Printf("Generated UUID %s\n", id)
	gnu.SayGnu()
}

