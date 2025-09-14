package main

import "fmt"

func main() {
	var operation string
	var number1, number2 int

	fmt.Println("calculator")
	fmt.Println("*******************")
	fmt.Println("Enter the operation you want? (add,subtract,multiply,divide)")
	fmt.Scanln(&operation)
	fmt.Println("Enter 1st number")
	fmt.Scanln(&number1)
	fmt.Println("Enter 2nd number")
	fmt.Scanln(&number2)

	switch operation {
	case "add":
		fmt.Printf("The sum of %d and %d is %d\n", number1, number2, number1+number2)
	case "subtract":
		fmt.Printf("The difference of %d and %d is %d\n", number1, number2, number1-number2)
	case "multiply":
		fmt.Printf("The product of %d and %d is %d\n", number1, number2, number1*number2)
	case "divide":
		fmt.Printf("The division of %d by %d is %d\n", number1, number2, number1/number2)
	default:
		fmt.Println("Invalid operation.")
	}
}
