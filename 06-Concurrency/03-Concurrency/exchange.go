package main

import "fmt"

//without temp variable
func main() {
	var num1, num2 int8
	fmt.Print("enter number1: ")
	fmt.Scanf("%d\n", &num1)
	fmt.Print("enter number2: ")
	fmt.Scanf("%d", &num2)

	// num1 = num1 ^ num2
	// num2 = num1 ^ num2
	// num1 = num1 ^ num2

	num1, num2 = num2, num1

	fmt.Println("number1: ", num1, "\t number2: ", num2)
}
