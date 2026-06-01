package structure

import (
	"fmt"
	"testing"
)

func TestSlicing(t *testing.T) {
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	mySlice := []int{10, 20, 30, 40, 50}

	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	subSlice := mySlice[1:3]
	fmt.Println(subSlice)

	mySlice = append(mySlice, 10)
	mySlice = append(mySlice, 20, 30)
	fmt.Println(mySlice)

	anotherSlice := myArray[:]
	anotherSlice = append(anotherSlice, 40,50)
	fmt.Println(anotherSlice)
}