package structure

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T){
	myMap := make(map[string]int)
	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 8

	fmt.Println("Apples:", myMap["apple"])

	myMap["banana"] = 12

	delete(myMap, "orange")

	for key, value := range myMap{
		fmt.Printf("%s -> %d\n", key, value)
	}

	val, ok := myMap["pear"]
	if ok {
		fmt.Println("Pear's value:", val)
	}else{
		fmt.Println("Pear not found inmap")
	}
}