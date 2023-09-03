package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string
	fmt.Println("Welcome to my calculator!")
	fmt.Print("Please enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Please enter operator (+, -, / or *): ")
	fmt.Scan(&operator)
	fmt.Print("Please enter second number: ")
	fmt.Scan(&num2)
	switch operator {
	case "+":
		fmt.Printf("Result: %.2f", num1+num2)
	case "-":
		fmt.Printf("Result: %.2f", num1-num2)
	case "/":
		fmt.Printf("Result: %.2f", num1/num2)
	case "*":
		fmt.Printf("Result: %.2f", num1*num2)
	}
}
