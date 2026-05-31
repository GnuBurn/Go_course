package condi

import "fmt"

func checkScore(score int){
	if score >= 70{
		fmt.Println("PASSED")
	} else{
		fmt.Println("FAILED")
	}
}

func checkGrade(score int){
	var grade string
	if score >= 80 {
		grade = "A"
	} else if score >= 70 {
		grade = "B"
	} else if score >= 60 {
		grade = "C"
	} else {
		grade = "F"
	}
	fmt.Println("Grade:", grade)
}

func checkDay(date int) string{
	switch date {
		case 1: return "Monday"
		case 2: return "Tuesday"
		case 3: return "Wednesday"
		case 4: return "Thursday"
		case 5: return "Friday"
		case 6: return "Saturday"
		case 7: return "Sunday"
		default: return "Invalid day"
	}
}

func IfTest(){
	checkScore(71)
	checkGrade(60)
	var day string = checkDay(3)
	fmt.Println(day)

	num1 := 5
	num2 := 10

	if sumNum := num1 + num2; sumNum >= 10 {
		fmt.Println("sumNum more than 10")
	}else{
    fmt.Println("sumNum less than 10")
  }
}