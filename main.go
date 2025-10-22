package main

import "fmt"

func addition(num1, num2 float32) float32 {
	return num1 + num2
}
func divition(num1, num2 float32) float32 {
	return num1 / num2
}
func substraction(num1, num2 float32) float32 {
	return num1 - num2
}
func multiplication(num1, num2 float32) float32 {
	return num1 * num2
}
func main() {

	var num1, num2, result float32
	var operator string

	fmt.Println("Enter your first number:")
	fmt.Scan(&num1)
	fmt.Println("Enter your second number:")
	fmt.Scan(&num2)
	fmt.Println("Select your operation within (+, -, *, /):")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		result = addition(num1, num2)
	case "-":
		result = substraction(num1, num2)
	case "*":
		result = multiplication(num1, num2)
	default:
		fmt.Println("Invalid operation")
		return
	}
	fmt.Printf("result= %v", result)
}
