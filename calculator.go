package main 

import (
	"fmt"
)
func main(){
	var number1, number2 float64
	var operator string 

	fmt.Print("Enter the first number :")
	fmt.Scanln(&number1)


	fmt.Print("Enter the second number :")
	fmt.Scanln(&number2)


	fmt.Println("Enter the Operators (+ - * / ) :")
	fmt.Scanln(&operator)

	switch operator {
	case "+" :
		fmt.Printf("%.3f %s %.3f = %f", number1, operator, number2, number1 + number2)
	case "-" :
		fmt.Printf("%.3f %s %.3f = %.3f", number1, operator, number2, number1 - number2)
	case "*" :
		fmt.Printf("%.3f %s %.3f = %.3f", number1, operator, number2, number1 * number2)
	case "/" :
		if number2 == 0.0 {
			fmt.Println("divide by zero situation")
		} else{
			fmt.Printf("%.3f %s %.3f = %.3f", number1, operator, number2, number1/number2)
		}
	default:
		fmt.Println("invalid operator")
	}
}