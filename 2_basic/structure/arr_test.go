package structure

import (
	"fmt"
	"testing"
)

type Person struct{
	Name string
	Age int
}

type Rectangle struct{
	Width, Height float64
}

type Shape interface{
	Area() float64
	Perimeter() float64
}

func (r Rectangle) Area() float64{
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64{
	return 2 * (r.Width + r.Height)
}

func showResult(s Shape) {
	fmt.Printf("Rectangle Area via Interface: %.2f\nRectangle Perimeter via Interface: %.2f\n", s.Area(), s.Perimeter())
}

func TestFunc(t *testing.T){
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	fmt.Println(myArray)

	for i := 0; i < len(myArray); i++{
		fmt.Println(myArray[i])
	}

	rect := Rectangle{Width: 10.0, Height: 5.0}
	showResult(rect)
}